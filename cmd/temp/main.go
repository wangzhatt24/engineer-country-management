package main

import (
	"context"
	"engineer-country-management/internal/pkg/mysql"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

type Country struct {
	Id          int64     `json:"id"`
	CountryName string    `json:"country_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main() {
	/*
		goi tat ca du lieu country trong db
		// them vao elastic search

	*/

	db := mysql.GetClient()

	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	esClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("\nerror when connect to elastic search")
	}

	_, err = esClient.Indices.Create("country").Do(context.TODO())
	if err != nil {
		log.Fatal("create index failed", err)
	}

	rows, err := db.Query("SELECT * FROM country")
	if err != nil {
		log.Fatal(err)
	}

	var country Country

	for rows.Next() {
		err := rows.Scan(&country.Id, &country.CountryName, &country.CreatedAt, &country.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}

		res, err :=
			esClient.
				Index("country").
				Id(fmt.Sprintf("%s:%d", "country", country.Id)).
				Request(country).
				Do(context.TODO())

		if err != nil {
			return
		}

		log.Printf(res.Result.Name)
	}

	log.Printf("all done")
}
