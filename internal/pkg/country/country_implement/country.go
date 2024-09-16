package country_implement

import (
	"context"
	"engineer-country-management/internal/pkg/country/country_elastic_search"
	dbHandler "engineer-country-management/internal/pkg/database_handlers/country"
	"engineer-country-management/internal/pkg/redis_cache"
	pb "engineer-country-management/pkg/country/v1"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type CountryImplement struct {
	pb.UnimplementedCountryServiceServer
	EsClient   *elasticsearch.TypedClient
	RedisCache *redis_cache.RedisCache
	DBHandler  *dbHandler.CountryDatabaseHandler
}

// method sections
func (ci *CountryImplement) GetCountryById(ctx context.Context, in *pb.GetCountryRequest) (*pb.Country, error) {
	country, err := ci.RedisCache.RedisFetchCountryById(ctx, in)

	if err != nil {

		country, err := ci.DBHandler.MysqlFetchCountryById(in.GetId())
		if err != nil {
			return nil, err
		}

		ci.RedisCache.RedisUpdateCountryById(ctx, country)
		return country, nil
	}

	return country, nil
}

func (ci *CountryImplement) AddCountry(ctx context.Context, in *pb.AddCountryRequest) (*pb.Country, error) {
	country, err := ci.DBHandler.MysqlAddCountry(ctx, in)

	if err != nil {
		return nil, err
	}

	return country, nil
}

// should be named DeleteCountryById
func (ci *CountryImplement) DeleteCountry(ctx context.Context, in *pb.DeleteCountryRequest) (*pb.Country, error) {
	country, err := ci.DBHandler.MysqlDeleteCountry(ctx, in)
	if err != nil {
		return nil, err
	}

	err = ci.RedisCache.RedisDeleteCountry(ctx, in)
	if err != nil {
		fmt.Printf("\n deleted in mysql but error orcur when deleting in redis")
	}

	return country, nil
}

func (ci *CountryImplement) UpdateCountry(ctx context.Context, in *pb.UpdateCountryRequest) (*pb.Country, error) {
	country, err := ci.DBHandler.MysqlUpdateCountry(ctx, in)
	if err != nil {
		return nil, err
	}

	err = ci.RedisCache.RedisUpdateCountryById(ctx, country)
	if err != nil {
		fmt.Println(err)
	}

	return country, nil
}

func (ci *CountryImplement) ListCountries(ctx context.Context, in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {
	listCountriesResponse, err := ci.DBHandler.MysqlListCountries(in)
	if err != nil {
		return nil, err
	}

	return listCountriesResponse, nil
}

func (ci *CountryImplement) SearchCountryFuzzyByName(ctx context.Context, in *pb.SearchCountryFuzzyByNameRequest) (*pb.Countries, error) {
	return country_elastic_search.SearchCountryFuzzyByName(ctx, ci.EsClient, in.GetCountryName())
}

// end method sections
