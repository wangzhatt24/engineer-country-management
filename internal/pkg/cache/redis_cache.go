package cache

import (
	"context"
	"engineer-country-management/internal/pkg/redis"
	pb "engineer-country-management/pkg/country/v1"
	"fmt"
	"time"

	"google.golang.org/protobuf/proto"
)

var redisClient = redis.GetClient()

// redis sections
func RedisGetCountryKey(id int64) string {
	return fmt.Sprintf("%s%d", "country:", id)
}

func RedisGetCountCountryKey(id int64) string {
	return fmt.Sprintf("count:country:%d", id)
}

func RedisFetchCountryById(ctx context.Context, in *pb.GetCountryRequest) (*pb.Country, error) {
	countryBytes, err := redisClient.Get(ctx, RedisGetCountryKey(in.GetId())).Bytes()
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

func RedisUpdateCountryById(ctx context.Context, country *pb.Country) error {
	countryBytes, err := proto.Marshal(country)
	if err != nil {
		return fmt.Errorf("\nconvert bytes error %v", err)
	}
	_, err = redisClient.Set(ctx, RedisGetCountryKey(country.GetId()), countryBytes, time.Hour).Result()
	if err != nil {
		return fmt.Errorf("\nerror when update redis %v", err)
	} else {
		fmt.Printf("\nupdated contry %v to redis", country.GetId())
		return nil
	}
}

func RedisDeleteCountry(ctx context.Context, in *pb.DeleteCountryRequest) error {
	_, err := redisClient.Del(ctx, RedisGetCountryKey(in.GetId())).Result()
	if err != nil {
		return err
	}

	return nil
}
