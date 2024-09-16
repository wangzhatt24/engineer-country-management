package main

import (
	country_queue "engineer-country-management/internal/pkg/queues/country/producer"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adjust/rmq/v5"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
)

func main() {

	// init connection
	errChan := make(chan<- error)
	connection, err := rmq.OpenConnection("country_count_producer_connection", "tcp", "localhost:6379", 0, errChan)
	if err != nil {
		panic(err)
	}

	// init queue
	countryCountQueue, err := connection.OpenQueue("country_count_queue")
	if err != nil {
		panic(err)
	}

	// init redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// init cron and publish job
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		countryQueue := country_queue.CountryCountProducer{
			RedisClient: redisClient,
		}

		err := countryQueue.CountryCountPublish(&connection, &countryCountQueue)
		if err != nil {
			log.Println(err)
		}
	})

	c.Start()

	// Listen to os sisnal (Ctrl+C, SIGINT, SIGTERM)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Producer stopped")
}
