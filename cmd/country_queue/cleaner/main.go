package main

import (
	"log"
	"time"

	"github.com/adjust/rmq/v5"
)

func main() {
	errChan := make(chan<- error)
	connection, err := rmq.OpenConnection("country_count_cleaner_connection", "tcp", "localhost:6379", 0, errChan)
	if err != nil {
		panic(err)
	}

	cleaner := rmq.NewCleaner(connection)

	for range time.Tick(time.Second) {
		returned, err := cleaner.Clean()
		if err != nil {
			log.Printf("failed to clean: %s", err)
			continue
		}
		log.Printf("cleaned %d", returned)
	}
}
