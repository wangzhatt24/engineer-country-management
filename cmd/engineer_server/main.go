package main

import (
	"database/sql"
	engineer_implement "engineer-country-management/internal/pkg/engineer/implement"
	pb "engineer-country-management/pkg/engineer/v1"
	"flag"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	db, err := sql.Open("mysql", "tyler:abc@123@tcp(127.0.0.1:3306)/engineer-country?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// init implement
	engineerImplement := engineer_implement.EngineerImplement{
		DB: db,
	}

	s := grpc.NewServer()
	pb.RegisterEngineerServiceServer(s, &engineerImplement)
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
