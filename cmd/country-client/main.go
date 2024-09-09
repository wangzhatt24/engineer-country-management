package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "engineer-country-management/pkg/country/v1"

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
	c := pb.NewCountryServiceClient(conn)

	// init context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// r, err := c.AddCountry(ctx, &pb.AddCountryRequest{CountryName: "New0"})
	r, err := c.GetCountryById(ctx, &pb.GetCountryRequest{Id: 110})

	// r, err := c.DeleteCountry(ctx, &pb.DeleteCountryRequest{Id: 247})
	// r, err := c.UpdateCountry(ctx, &pb.UpdateCountryRequest{Id: 230, CountryName: "New 230"})
	// r, err := c.ListCountries(ctx, &pb.ListCountriesRequest{
	// 	PageSize: 10, PageNumber: 3,
	// })

	if err != nil {
		log.Fatal("error when get country by id", err)
	}

	log.Printf("Result: %s", r)
}
