package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	pb "engineer-country-management/pkg/country/v1"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedCountryServiceServer
	db          *sql.DB
	redisClient *redis.Client
}

// redis sections
func redisGetCountryKey(id int64) string {
	return fmt.Sprintf("%s%d", "country:", id)
}

func (s *server) redisFetchCountryById(ctx context.Context, in *pb.GetCountryRequest) (*pb.Country, error) {
	countryBytes, err := s.redisClient.Get(ctx, redisGetCountryKey(in.GetId())).Bytes()
	// loi doc cache
	// tra ve nil

	if err != nil {
		fmt.Printf("\nerror when reading redis (GetCountryById): %v", err)
		return nil, fmt.Errorf("\nerror when reading redis (GetCountryById): %v", err)
	}

	if countryBytes == nil {
		fmt.Printf("\ncountry not found in redis %v (GetCountryById)", in.GetId())
		return nil, fmt.Errorf("\ncountry not found in redis %v (GetCountryById)", in.GetId())
	} else {
		var country pb.Country
		if err := proto.Unmarshal(countryBytes, &country); err != nil {
			return nil, fmt.Errorf("\nerror when unmarshal country bytes")
		} else {
			return &country, nil
		}
	}
}

func (s *server) redisUpdateCountryById(ctx context.Context, country *pb.Country) error {
	countryBytes, err := proto.Marshal(country)
	if err != nil {
		return fmt.Errorf("\nconvert bytes error %v", err)
	}
	_, err = s.redisClient.Set(ctx, redisGetCountryKey(country.GetId()), countryBytes, time.Hour).Result()
	if err != nil {
		return fmt.Errorf("\nerror when update redis\n%v", err)
	} else {
		fmt.Printf("\nupdated contry %v to redis", country.GetId())
		return nil
	}
}

// end redis sections

// mysql sections
func (s *server) mysqlFetchCountryById(in *pb.GetCountryRequest) (*pb.Country, error) {
	row := s.db.QueryRow("SELECT * FROM country WHERE id = ?", in.GetId())

	var country pb.Country
	var created_at time.Time
	var updated_at time.Time

	err := row.Scan(&country.Id, &country.CountryName, &created_at, &updated_at)
	if err != nil {
		return nil, err
	}

	return &pb.Country{
		Id:          country.Id,
		CountryName: country.CountryName,
		CreatedAt:   timestamppb.New(created_at),
		UpdatedAt:   timestamppb.New(updated_at),
	}, nil
}

func (s *server) mysqlAddCountry(in *pb.AddCountryRequest) (*pb.Country, error) {
	created_at, updated_at := time.Now(), time.Now()
	result, err := s.db.Exec("INSERT INTO country(country_name, created_at, updated_at) VALUES (?, ?, ?)", in.GetCountryName(), created_at, updated_at)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &pb.Country{
		Id:          id,
		CountryName: in.CountryName,
		CreatedAt:   timestamppb.New(created_at),
		UpdatedAt:   timestamppb.New(updated_at),
	}, nil
}

// end mysql sections

// method sections
func (s *server) GetCountryById(ctx context.Context, in *pb.GetCountryRequest) (*pb.Country, error) {
	country, err := s.redisFetchCountryById(ctx, in)
	if err != nil {
		country, err := s.mysqlFetchCountryById(in)
		if err != nil {
			return nil, err
		}

		return country, err
	}

	return country, err
}

func (s *server) AddCountry(ctx context.Context, in *pb.AddCountryRequest) (*pb.Country, error) {
	country, err := s.mysqlAddCountry(in)

	if err != nil {
		return nil, err
	}

	return country, nil
}

// should be named DeleteCountryById
func (s *server) DeleteCountry(ctx context.Context, in *pb.DeleteCountryRequest) (*pb.Country, error) {
	country, err := s.GetCountryById(ctx, &pb.GetCountryRequest{Id: in.Id})
	if err != nil {
		return nil, errors.New("error when deleting country")
	}

	result, err := s.db.Exec("DELETE FROM country where id = ?", in.Id)

	if err != nil {
		log.Fatalf("error when delete country %v", err)
	}

	ra, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("error checking row affected %v", ra)
	}

	if ra == 1 {
		return country, nil
	}

	return nil, errors.New("error when deleting country")
}

func (s *server) UpdateCountry(ctx context.Context, in *pb.UpdateCountryRequest) (*pb.Country, error) {
	result, err := s.db.Exec("UPDATE country SET country_name = ? WHERE country.id = ?", in.GetCountryName(), in.GetId())

	if err != nil {
		return nil, err
	}

	rf, err := result.RowsAffected()
	if err != nil {
		fmt.Println("\nerror when check row affected")
	}

	if rf != 1 {
		fmt.Println("\nno rows updated")
	}

	country, err := s.mysqlFetchCountryById(&pb.GetCountryRequest{Id: in.GetId()})
	if err != nil {
		return nil, fmt.Errorf("error when fetch country by id (mysqlFetchCountryById): %v", err)
	}

	err = s.redisUpdateCountryById(ctx, country)
	if err != nil {
		fmt.Printf("\nerror when update country to redis %v", err)
	}

	return country, nil
}

func (s *server) ListCountries(ctx context.Context, in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {
	// get total count
	// get records
	pageNumber := in.GetPageNumber()
	pageSize := in.GetPageSize()

	// ensure that page number is valid
	if pageNumber <= 0 {
		pageNumber = 1 // default
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (pageNumber - 1) * pageSize

	// total countries
	var totalCount int64
	err := s.db.QueryRow("SELECT COUNT(*) FROM country").Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query("SELECT * FROM country LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, err
	}

	var countries []*pb.Country

	for rows.Next() {
		var country pb.Country
		var created_at time.Time
		// var updated_at time.Time

		err := rows.Scan(&country.Id, &country.CountryName, &created_at, &created_at)
		if err != nil {
			return nil, err
		}

		countries = append(countries, &country)
	}

	return &pb.ListCountriesResponse{
		Countries:  countries,
		TotalCount: totalCount,
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}, nil
}

// end method sections

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	db, err := sql.Open("mysql", "tyler:abc@123@tcp(127.0.0.1:3306)/engineer-country?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCountryServiceServer(s, &server{db: db, redisClient: redisClient})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
