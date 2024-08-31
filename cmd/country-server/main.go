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

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedCountryServiceServer
	db *sql.DB
}

func (s server) AddCountry(ctx context.Context, in *pb.AddCountryRequest) (*pb.Country, error) {
	created_at, updated_at := time.Now(), time.Now()
	result, err := s.db.Exec("INSERT INTO country(country_name, created_at, updated_at) VALUES (?, ?, ?)", in.CountryName, created_at, updated_at)

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

func (s server) GetCountryById(ctx context.Context, in *pb.GetCountryRequest) (*pb.Country, error) {
	row := s.db.QueryRow("SELECT * from country where id = ?", in.Id)

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

// should be named DeleteCountryById
func (s server) DeleteCountry(ctx context.Context, in *pb.DeleteCountryRequest) (*pb.Country, error) {
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

func (s server) UpdateCountry(ctx context.Context, in *pb.UpdateCountryRequest) (*pb.Country, error) {
	result, err := s.db.Exec("UPDATE country SET country_name = ? WHERE country.id = ?", in.CountryName, in.Id)

	if err != nil {
		return nil, err
	}

	rf, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rf != 1 {
		return nil, errors.New("no country updated")
	}

	country, err := s.GetCountryById(ctx, &pb.GetCountryRequest{Id: in.Id})
	if err != nil {
		return nil, err
	}

	return country, nil
}

func (s server) ListCountries(context.Context, *emptypb.Empty) (*pb.Countries, error) {
	result, err := s.db.Query("SELECT * FROM country")
	if err != nil {
		return nil, err
	}

	var countries []*pb.Country

	for result.Next() {
		var country pb.Country
		var created_at time.Time
		var updated_at time.Time

		err := result.Scan(&country.Id, &country.CountryName, &created_at, &updated_at)
		if err != nil {
			return nil, err
		}

		// time convert
		country.CreatedAt = timestamppb.New(created_at)
		country.UpdatedAt = timestamppb.New(updated_at)

		countries = append(countries, &country)
	}

	return &pb.Countries{Countries: countries}, nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
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
	pb.RegisterCountryServiceServer(s, &server{db: db})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
