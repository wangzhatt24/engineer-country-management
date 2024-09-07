package main

import (
	"context"
	"engineer-country-management/internal/pkg/redis"
	"fmt"
)

func main() {
	redisClient := redis.GetClient()

	pPub := redisClient.PSubscribe(context.Background(), "log:count:country:*")
	defer pPub.Close()

	ch := pPub.Channel()

	for msg := range ch {
		fmt.Println(msg)
	}
}
