package countryimplement

import (
	"context"
	"engineer-country-management/internal/pkg/country_elastic_search"
	dbHandler "engineer-country-management/internal/pkg/database_handlers/country"
	"engineer-country-management/internal/pkg/redis_cache"
	pb "engineer-country-management/pkg/country/v1"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type Server struct {
	pb.UnimplementedCountryServiceServer
	EsClient   *elasticsearch.TypedClient
	RedisCache *redis_cache.RedisCache
	DBHandler  *dbHandler.CountryDatabaseHandler
}

// method sections
func (s *Server) GetCountryById(ctx context.Context, in *pb.GetCountryRequest) (*pb.Country, error) {
	country, err := s.RedisCache.RedisFetchCountryById(ctx, in)

	if err != nil {

		country, err := s.DBHandler.MysqlFetchCountryById(in.GetId())
		if err != nil {
			return nil, err
		}

		s.RedisCache.RedisUpdateCountryById(ctx, country)
		return country, nil
	}

	return country, nil
}

func (s *Server) AddCountry(ctx context.Context, in *pb.AddCountryRequest) (*pb.Country, error) {
	country, err := s.DBHandler.MysqlAddCountry(ctx, in)

	if err != nil {
		return nil, err
	}

	return country, nil
}

// should be named DeleteCountryById
func (s *Server) DeleteCountry(ctx context.Context, in *pb.DeleteCountryRequest) (*pb.Country, error) {
	country, err := s.DBHandler.MysqlDeleteCountry(ctx, in)
	if err != nil {
		return nil, err
	}

	err = s.RedisCache.RedisDeleteCountry(ctx, in)
	if err != nil {
		fmt.Printf("\n deleted in mysql but error orcur when deleting in redis")
	}

	return country, nil
}

func (s *Server) UpdateCountry(ctx context.Context, in *pb.UpdateCountryRequest) (*pb.Country, error) {
	country, err := s.DBHandler.MysqlUpdateCountry(ctx, in)
	if err != nil {
		return nil, err
	}

	err = s.RedisCache.RedisUpdateCountryById(ctx, country)
	if err != nil {
		fmt.Println(err)
	}

	return country, nil
}

func (s *Server) ListCountries(ctx context.Context, in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {
	listCountriesResponse, err := s.DBHandler.MysqlListCountries(in)
	if err != nil {
		return nil, err
	}

	return listCountriesResponse, nil
}

func (s *Server) SearchCountryFuzzyByName(ctx context.Context, in *pb.SearchCountryFuzzyByNameRequest) (*pb.Countries, error) {
	return country_elastic_search.SearchCountryFuzzyByName(ctx, s.EsClient, in.GetCountryName())
}

// end method sections
