package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"engineer-country-management/cmd/country-server/internal/interceptors"
	"engineer-country-management/internal/pkg/cache"
	redisWrapper "engineer-country-management/internal/pkg/redis"
	pb "engineer-country-management/pkg/country/v1"

	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedCountryServiceServer
	db          *sql.DB
	redisClient *redis.Client
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

func (s *server) mysqlDeleteCountry(ctx context.Context, in *pb.DeleteCountryRequest) (*pb.Country, error) {
	country, err := s.GetCountryById(ctx, &pb.GetCountryRequest{Id: in.Id})
	if err != nil {
		return nil, fmt.Errorf("\nerror when deleting country (mysqlDeleteCountry) %v", err)
	}

	result, err := s.db.Exec("DELETE FROM country where id = ?", in.Id)

	if err != nil {
		return nil, fmt.Errorf("\nerror when delete country %v", err)
	}

	ra, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("\nerror checking row affected %v", err)
	}

	if ra == 1 {
		return country, nil
	}

	return nil, fmt.Errorf("\nerror when deleting country")
}

func (s *server) mysqlUpdateCountry(ctx context.Context, in *pb.UpdateCountryRequest) (*pb.Country, error) {
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

	err = cache.RedisUpdateCountryById(ctx, country)
	if err != nil {
		fmt.Printf("\nerror when update country to redis %v", err)
	}

	return country, nil
}

func (s *server) mysqlListCountries(in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {
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

	var countries pb.Countries

	for rows.Next() {
		var country pb.Country
		var created_at time.Time
		// var updated_at time.Time

		err := rows.Scan(&country.Id, &country.CountryName, &created_at, &created_at)
		if err != nil {
			return nil, err
		}

		countries.Country = append(countries.Country, &country)
	}

	return &pb.ListCountriesResponse{
		Countries:  &countries,
		TotalCount: totalCount,
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}, nil
}

// end mysql sections

// method sections
func (s *server) GetCountryById(ctx context.Context, in *pb.GetCountryRequest) (*pb.Country, error) {
	country, err := cache.RedisFetchCountryById(ctx, in)
	if err != nil {
		country, err := s.mysqlFetchCountryById(in)
		if err != nil {
			return nil, err
		}

		cache.RedisUpdateCountryById(ctx, country)
		return country, nil
	}

	return country, nil
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
	country, err := s.mysqlDeleteCountry(ctx, in)
	if err != nil {
		return nil, err
	}

	err = cache.RedisDeleteCountry(ctx, in)
	if err != nil {
		fmt.Printf("\n deleted in mysql but error orcur when deleting in redis")
	}

	return country, nil
}

func (s *server) UpdateCountry(ctx context.Context, in *pb.UpdateCountryRequest) (*pb.Country, error) {
	country, err := s.mysqlUpdateCountry(ctx, in)
	if err != nil {
		return nil, err
	}

	err = cache.RedisUpdateCountryById(ctx, country)
	if err != nil {
		fmt.Println(err)
	}

	return country, nil
}

func (s *server) ListCountries(ctx context.Context, in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {
	listCountriesResponse, err := s.mysqlListCountries(in)
	if err != nil {
		return nil, err
	}

	return listCountriesResponse, nil
}

// end method sections

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	db, err := sql.Open("mysql", "tyler:abc@123@tcp(127.0.0.1:3306)/engineer-country?parseTime=true")
	redisClient := redisWrapper.GetClient()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.IncrementCountryCountInterceptor),
	)
	pb.RegisterCountryServiceServer(s, &server{db: db, redisClient: redisClient})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
