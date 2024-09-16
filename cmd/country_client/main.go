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

	now := time.Now()

	// r, err := c.AddCountry(ctx, &pb.AddCountryRequest{CountryName: "Tyler"})
	// r, err := c.GetCountryById(ctx, &pb.GetCountryRequest{Id: 304})

	r, err := c.DeleteCountry(ctx, &pb.DeleteCountryRequest{Id: 304})
	// r, err := c.UpdateCountry(ctx, &pb.UpdateCountryRequest{Id: 304, CountryName: "New 2"})
	// r, err := c.ListCountries(ctx, &pb.ListCountriesRequest{
	// 	PageSize: 10, PageNumber: 3,
	// })

	// r, err := c.SearchCountryFuzzyByName(ctx, &pb.SearchCountryFuzzyByNameRequest{
	// 	CountryName: "vi",
	// })

	duration := time.Since(now)
	log.Println(duration)

	if err != nil {
		log.Fatal("error when get country by id", err)
	}

	log.Printf("Result: %s", r)
}
