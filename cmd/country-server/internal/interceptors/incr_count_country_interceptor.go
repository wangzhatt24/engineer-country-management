package interceptors

import (
	"context"
	"engineer-country-management/internal/pkg/cache"
	redisWrapper "engineer-country-management/internal/pkg/redis"
	pb "engineer-country-management/pkg/country/v1"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func IncrementCountryCountInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Call the actual RPC handler
	resp, err := handler(ctx, req)

	// Check if the request is for GetCountryById and if the call was successful
	if err == nil && info.FullMethod == "/country.v1.CountryService/GetCountryById" {
		// Assuming the request contains a country ID, cast it and call Redis INCR
		if countryRequest, ok := req.(*pb.GetCountryRequest); ok {
			redisClient := redisWrapper.GetClient()
			countryID := countryRequest.GetId() // Lấy ID quốc gia
			redisKey := cache.RedisGetCountCountryKey(countryID)
			channel := fmt.Sprintf("log:count:country:%d", countryID)

			// Increase count in Redis asynchronously
			go func() {
				err := redisClient.Incr(context.Background(), redisKey).Err()
				if err != nil {
					log.Printf("Error incrementing country count for ID %d: %v", countryID, err)
				} else {
					redisClient.Publish(context.Background(), channel, fmt.Sprintf("%d just incre", countryID))
					log.Printf("Successfully incremented count for country ID: %d", countryID)
				}
			}()
		}
	} else {
		log.Printf("Cant catch method %v", err)
	}

	return resp, err
}
