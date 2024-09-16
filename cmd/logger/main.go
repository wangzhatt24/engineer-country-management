package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// init redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pPub := redisClient.PSubscribe(context.Background(), "log:count:country:*")
	defer pPub.Close()

	ch := pPub.Channel()

	for msg := range ch {
		fmt.Println(msg)
	}
}
