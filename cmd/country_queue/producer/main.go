package main

import (
	country_queue "engineer-country-management/internal/pkg/queues/country/producer"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adjust/rmq/v5"
	"github.com/robfig/cron/v3"
)

func main() {
	errChan := make(chan<- error)
	connection, err := rmq.OpenConnection("country_count_producer_connection", "tcp", "localhost:6379", 0, errChan)
	if err != nil {
		panic(err)
	}

	countryCountQueue, err := connection.OpenQueue("country_count_queue")
	if err != nil {
		panic(err)
	}

	c := cron.New()
	c.AddFunc("@every 1s", func() {
		err := country_queue.CountryCountPublish(&connection, &countryCountQueue)
		if err != nil {
			log.Println(err)
		}
	})

	c.Start()

	// Bắt tín hiệu từ hệ điều hành (Ctrl+C, SIGINT, SIGTERM)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Chờ tín hiệu để dừng chương trình
	<-sigs

	log.Println("Shutting down producer...")

	log.Println("Producer stopped")
}
