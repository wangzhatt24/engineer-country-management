package main

import (
	consumer "engineer-country-management/internal/pkg/queues/elastic_search/consumer"
	QUEUE_CONSTANTS "engineer-country-management/internal/pkg/queues/elastic_search/elastic_search_queue_constants"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adjust/rmq/v5"
)

func main() {
	errChan := make(chan error, 10)
	go logErrors(errChan)

	connection, err := rmq.OpenConnection(QUEUE_CONSTANTS.CONSUNMER_CONNECTION_NAME, "tcp", "localhost:6379", 0, errChan)
	if err != nil {
		log.Fatal(err)
	}

	elasticSearchCountryQueue, err := connection.OpenQueue(QUEUE_CONSTANTS.QUEUE_NAME)
	if err != nil {
		panic(err)
	}

	// consumer
	err = elasticSearchCountryQueue.StartConsuming(10, time.Second)
	if err != nil {
		log.Fatalf("\nerror when start consuming %v", err)
	}

	elasticSearchCountryConsumer := &consumer.QueueElasticSearchConsumer{}
	name, err := elasticSearchCountryQueue.AddConsumer(QUEUE_CONSTANTS.CONSUMER_NAME, elasticSearchCountryConsumer)
	if err != nil {
		log.Fatalf("\nadd consumer failed %v", err)
	}

	fmt.Printf("\nnew consumer: %v", name)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signals)

	<-signals // wait for signal
	go func() {
		<-signals // hard exit on second signal (in case shutdown gets stuck)
		os.Exit(1)
	}()

	<-connection.StopAllConsuming() // wait for all Consume() calls to finish
}

func logErrors(errChan <-chan error) {
	for err := range errChan {
		switch err := err.(type) {
		case *rmq.HeartbeatError:
			if err.Count == rmq.HeartbeatErrorLimit {
				log.Print("heartbeat error (limit): ", err)
			} else {
				log.Print("heartbeat error: ", err)
			}
		case *rmq.ConsumeError:
			log.Print("consume error: ", err)
		case *rmq.DeliveryError:
			log.Print("delivery error: ", err.Delivery, err)
		default:
			log.Print("other error: ", err)
		}
	}
}
