package country_interceptors

import (
	"context"
	queue_elastic_search_producer "engineer-country-management/internal/pkg/queues/elastic_search/producer"
	"engineer-country-management/internal/pkg/redis_cache"
	pb "engineer-country-management/pkg/country/v1"
	"fmt"
	"log"

	"github.com/adjust/rmq/v5"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func GetCountryInterceptor(redisClient *redis.Client, queue *rmq.Queue) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		// Gọi handler để xử lý yêu cầu
		resp, err := handler(ctx, req)

		// Nếu không có lỗi, kiểm tra phương thức GRPC và xử lý tương ứng
		if err == nil {
			switch info.FullMethod {
			case "/country.v1.CountryService/GetCountryById":
				handleCountCountry(&resp, redisClient) // Truyền redisClient vào
			case "/country.v1.CountryService/AddCountry":
				// handleAddCountry(&resp)
			case "/country.v1.CountryService/UpdateCountry":
				// handleUpdateCountry(&resp)
			case "/country.v1.CountryService/DeleteCountry":
				handleDeleteCountry(&resp, queue)
			default:
				log.Println("no interceptor handler called", info.FullMethod)
			}
		} else {
			log.Printf("error in method %v", err)
		}

		return resp, err
	}
}

func handleDeleteCountry(resp *interface{}, queue *rmq.Queue) {
	log.Println("handle delete country interceptor")

	country, ok := (*resp).(*pb.Country)
	if !ok {
		log.Println("err when convert country from request")
	}

	producer := queue_elastic_search_producer.QueueElasticSearchProducer{
		Queue: queue,
	}

	err := producer.DeleteCountryPublish(queue_elastic_search_producer.Country{
		Id:          country.GetId(),
		CountryName: country.GetCountryName(),
		CreatedAt:   country.GetCreatedAt().AsTime(),
		UpdatedAt:   country.UpdatedAt.AsTime(),
	})

	if err != nil {
		log.Printf("\nhandleDeleteCountryError: %v", err)
	}
}

// func handleUpdateCountry(resp *interface{}) {
// 	log.Printf("\nhandle update country interceptor")

// 	country, ok := (*resp).(*pb.Country)
// 	if !ok {
// 		log.Printf("\nerr when convert country from request %v", ok)
// 	}

// 	errDelete := queue_elastic_search_producer.UpdateCountryPublish(conn, queue, queue_elastic_search_producer.Country{
// 		Id:          country.GetId(),
// 		CountryName: country.GetCountryName(),
// 		CreatedAt:   country.GetCreatedAt().AsTime(),
// 		UpdatedAt:   country.UpdatedAt.AsTime(),
// 	})

// 	if errDelete != nil {
// 		log.Println(errDelete)
// 	}
// }

// func handleAddCountry(resp *interface{}) {
// 	log.Printf("handle add country interceptor")

// 	country, ok := (*resp).(*pb.Country)
// 	if !ok {
// 		log.Printf("\nerr when convert country from request %v", ok)
// 	}

// 	queue_elastic_search_producer.AddCountryPublish(conn, queue, queue_elastic_search_producer.Country{
// 		Id:          country.GetId(),
// 		CountryName: country.GetCountryName(),
// 		CreatedAt:   country.GetCreatedAt().AsTime(),
// 		UpdatedAt:   country.UpdatedAt.AsTime(),
// 	})
// }

func handleCountCountry(resp *interface{}, redisClient *redis.Client) {
	if country, ok := (*resp).(*pb.Country); ok {
		countryId := country.GetId()
		redisKey := redis_cache.RedisGetCountCountryKey(countryId)
		channel := fmt.Sprintf("log:count:country:%d", countryId)

		// Increase count in Redis asynchronously
		go func() {
			err := redisClient.Incr(context.Background(), redisKey).Err()
			if err != nil {
				log.Printf("error incrementing country count for ID %d: %v", countryId, err)
			} else {
				// for log service
				redisClient.Publish(context.Background(), channel, fmt.Sprintf("%d just incre", countryId))
				log.Printf("successfully incremented count for country ID: %d", countryId)
			}
		}()
	} else {
		log.Println("request doenst contains data that match country")
	}
}
