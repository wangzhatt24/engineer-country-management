package main

import (
	"context"
	pb "engineer-country-management/pkg/engineer/v1"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	// set up connection
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("connect error", err)
	}

	defer conn.Close()

	// create client
	c := pb.NewEngineerServiceClient(conn)

	// init context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	now := time.Now()

	// r, err := c.GetEngineerById(ctx, &pb.GetEngineerByIdRequest{Id: 1})
	// r, err := c.AddEngineer(ctx,
	// 	&pb.AddEngineerRequest{
	// 		FirstName: "Tyler",
	// 		LastName:  "Wang",
	// 		Gender:    1,
	// 		CountryId: 231,
	// 		Title:     "Backend Developer",
	// 	},
	// )

	r, err := c.DeleteEngineerById(ctx, &pb.DeleteEngineerRequest{Id: 100001})

	// r, err := c.UpdateEngineer(ctx, &pb.UpdateEngineerRequest{
	// 	Id:        100003,
	// 	FirstName: "none",
	// 	LastName:  "wang",
	// 	Gender:    0,
	// 	Title:     "Golang Developer",
	// 	CountryId: 2,
	// })

	// r, err := c.ListEngineers(ctx, &pb.ListEngineersRequest{
	// 	PageNumber: 1, PageSize: 10,
	// })

	duration := time.Since(now)
	log.Print(duration)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Result: %s", r)
}
