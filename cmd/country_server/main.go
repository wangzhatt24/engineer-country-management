package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"

	"engineer-country-management/cmd/country_server/internal/country_interceptors"
	countryImplement "engineer-country-management/internal/pkg/country_implement"
	"engineer-country-management/internal/pkg/database_handlers/country"
	"engineer-country-management/internal/pkg/redis_cache"

	pb "engineer-country-management/pkg/country/v1"

	"github.com/adjust/rmq/v5"
	"github.com/elastic/go-elasticsearch/v8"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	// init listen port
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// init elastic search
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	esClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("\nerror when connect to elastic search")
	}

	// init redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// init redis cache - an internal wrapper for country server
	redisCacheCountry := &redis_cache.RedisCache{
		RedisClient: redisClient,
	}

	// init queue, but not handle errChan
	errChan := make(chan<- error)
	connection, err := rmq.OpenConnection("country_producer_connection", "tcp", "localhost:6379", 0, errChan)
	if err != nil {
		panic(err)
	}

	countryQueue, err := connection.OpenQueue("country_count_queue")
	if err != nil {
		panic(err)
	}

	// init connection to mysql
	dsn := "tyler:abc@123@tcp(127.0.0.1:3306)/engineer-country?parseTime=true"
	mysqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to connect to database: %w", err)
	}

	// init DBHandler
	redisCache := redis_cache.RedisCache{
		RedisClient: redisClient,
	}

	dbHandler := country.CountryDatabaseHandler{
		DB:         mysqlDB,
		RedisCache: &redisCache,
	}

	// init interceptor
	countryInterceptors := country_interceptors.GetCountryInterceptor(redisClient, &countryQueue)

	// init grpc server
	server := grpc.NewServer(
		grpc.UnaryInterceptor(countryInterceptors),
	)
	pb.RegisterCountryServiceServer(server, &countryImplement.Server{EsClient: esClient, RedisCache: redisCacheCountry, DBHandler: &dbHandler})
	log.Printf("server listening at %v", lis.Addr())

	// start service
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
