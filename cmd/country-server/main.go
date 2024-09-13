package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"

	"engineer-country-management/cmd/country-server/internal/country_interceptors"
	countryImplement "engineer-country-management/internal/pkg/country_implement"

	pb "engineer-country-management/pkg/country/v1"

	"google.golang.org/grpc"
)

// end redis sections

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(country_interceptors.IncrementCountryCountInterceptor),
	)
	pb.RegisterCountryServiceServer(s, &countryImplement.Server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
