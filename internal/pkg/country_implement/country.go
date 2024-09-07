package countryimplement

import (
	"context"
	"engineer-country-management/internal/pkg/cache"
	dbHandler "engineer-country-management/internal/pkg/database_handlers/country"
	pb "engineer-country-management/pkg/country/v1"
	"fmt"
)

type Server struct {
	pb.UnimplementedCountryServiceServer
}

// method sections
func (s *Server) GetCountryById(ctx context.Context, in *pb.GetCountryRequest) (*pb.Country, error) {
	country, err := cache.RedisFetchCountryById(ctx, in)
	if err != nil {
		// country, err := s.mysqlFetchCountryById(in)
		country, err := dbHandler.MysqlFetchCountryById(in.GetId())
		if err != nil {
			return nil, err
		}

		cache.RedisUpdateCountryById(ctx, country)
		return country, nil
	}

	return country, nil
}

func (s *Server) AddCountry(ctx context.Context, in *pb.AddCountryRequest) (*pb.Country, error) {
	country, err := dbHandler.MysqlAddCountry(ctx, in)

	if err != nil {
		return nil, err
	}

	return country, nil
}

// should be named DeleteCountryById
func (s *Server) DeleteCountry(ctx context.Context, in *pb.DeleteCountryRequest) (*pb.Country, error) {
	country, err := dbHandler.MysqlDeleteCountry(ctx, in)
	if err != nil {
		return nil, err
	}

	err = cache.RedisDeleteCountry(ctx, in)
	if err != nil {
		fmt.Printf("\n deleted in mysql but error orcur when deleting in redis")
	}

	return country, nil
}

func (s *Server) UpdateCountry(ctx context.Context, in *pb.UpdateCountryRequest) (*pb.Country, error) {
	country, err := dbHandler.MysqlUpdateCountry(ctx, in)
	if err != nil {
		return nil, err
	}

	err = cache.RedisUpdateCountryById(ctx, country)
	if err != nil {
		fmt.Println(err)
	}

	return country, nil
}

func (s *Server) ListCountries(ctx context.Context, in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {
	listCountriesResponse, err := dbHandler.MysqlListCountries(in)
	if err != nil {
		return nil, err
	}

	return listCountriesResponse, nil
}

// end method sections
