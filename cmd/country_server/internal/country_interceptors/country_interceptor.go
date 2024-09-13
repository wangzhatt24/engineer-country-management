package country_interceptors

import (
	"context"
	"engineer-country-management/internal/pkg/cache"
	"engineer-country-management/internal/pkg/queues/elastic_search/elastic_search_queue_constants"
	queue_elastic_search_producer "engineer-country-management/internal/pkg/queues/elastic_search/producer"
	redisWrapper "engineer-country-management/internal/pkg/redis"
	pb "engineer-country-management/pkg/country/v1"
	"fmt"
	"log"

	"github.com/adjust/rmq/v5"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func IncrementCountryCountInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	redisClient := redisWrapper.GetClient()

	errChan := make(chan<- error)
	esProdConn, err := rmq.OpenConnection(elastic_search_queue_constants.PRODUCER_CONNECTION_NAME, "tcp", "localhost:6379", 0, errChan)
	if err != nil {
		panic(err)
	}

	esCountryQueue, err := esProdConn.OpenQueue(elastic_search_queue_constants.QUEUE_NAME)
	if err != nil {
		panic(err)
	}

	// Call the actual RPC handler
	// This line will call the handler and get the respose
	// base on the response we will handle next
	resp, err := handler(ctx, req)

	// Check if the request is for GetCountryById and if the call was successful
	if err == nil {
		switch info.FullMethod {
		case "/country.v1.CountryService/GetCountryById":
			handleCountCountry(&resp, redisClient)
		case "/country.v1.CountryService/AddCountry":
			handleAddCountry(&resp, &esProdConn, &esCountryQueue)
		case "/country.v1.CountryService/UpdateCountry":
			handleUpdateCountry(&resp, &esProdConn, &esCountryQueue)
		case "/country.v1.CountryService/DeleteCountry":
			handleDeleteCountry(&resp, &esProdConn, &esCountryQueue)
		default:
			log.Println("no interceptor handler called", info.FullMethod)
		}

	} else {
		log.Printf("error in method %v", err)
	}

	return resp, err
}

func handleDeleteCountry(resp *interface{}, conn *rmq.Connection, queue *rmq.Queue) {
	log.Println("handle delete country interceptor")

	country, ok := (*resp).(*pb.Country)
	if !ok {
		log.Println("err when convert country from request")
	}

	errDelete := queue_elastic_search_producer.DeleteCountryPublish(conn, queue, queue_elastic_search_producer.Country{
		Id:          country.GetId(),
		CountryName: country.GetCountryName(),
		CreatedAt:   country.GetCreatedAt().AsTime(),
		UpdatedAt:   country.UpdatedAt.AsTime(),
	})

	if errDelete != nil {
		log.Printf("\nhandleDeleteCountryError: %v", errDelete)
	}
}

func handleUpdateCountry(resp *interface{}, conn *rmq.Connection, queue *rmq.Queue) {
	log.Printf("\nhandle update country interceptor")

	country, ok := (*resp).(*pb.Country)
	if !ok {
		log.Printf("\nerr when convert country from request %v", ok)
	}

	errDelete := queue_elastic_search_producer.UpdateCountryPublish(conn, queue, queue_elastic_search_producer.Country{
		Id:          country.GetId(),
		CountryName: country.GetCountryName(),
		CreatedAt:   country.GetCreatedAt().AsTime(),
		UpdatedAt:   country.UpdatedAt.AsTime(),
	})

	if errDelete != nil {
		log.Println(errDelete)
	}
}

func handleAddCountry(resp *interface{}, conn *rmq.Connection, queue *rmq.Queue) {
	log.Printf("handle add country interceptor")

	country, ok := (*resp).(*pb.Country)
	if !ok {
		log.Printf("\nerr when convert country from request %v", ok)
	}

	queue_elastic_search_producer.AddCountryPublish(conn, queue, queue_elastic_search_producer.Country{
		Id:          country.GetId(),
		CountryName: country.GetCountryName(),
		CreatedAt:   country.GetCreatedAt().AsTime(),
		UpdatedAt:   country.UpdatedAt.AsTime(),
	})
}

func handleCountCountry(req *interface{}, redisClient *redis.Client) {
	// Assuming the request contains a country ID, cast it and call Redis INCR
	if countryRequest, ok := (*req).(*pb.GetCountryRequest); ok {
		countryID := countryRequest.GetId() // Lấy ID quốc gia
		redisKey := cache.RedisGetCountCountryKey(countryID)
		channel := fmt.Sprintf("log:count:country:%d", countryID)

		// Increase count in Redis asynchronously
		go func() {
			err := redisClient.Incr(context.Background(), redisKey).Err()
			if err != nil {
				log.Printf("error incrementing country count for ID %d: %v", countryID, err)
			} else {
				// for log service
				redisClient.Publish(context.Background(), channel, fmt.Sprintf("%d just incre", countryID))
				log.Printf("successfully incremented count for country ID: %d", countryID)
			}
		}()
	} else {
		log.Println("request doenst contains data that doesnt match country")
	}
}
