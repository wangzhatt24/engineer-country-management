package queue_elastic_search_consumer

import (
	"context"
	"encoding/json"
	"engineer-country-management/internal/pkg/queues/elastic_search/elastic_search_queue_constants"
	producer "engineer-country-management/internal/pkg/queues/elastic_search/producer"
	"fmt"
	"log"

	"github.com/adjust/rmq/v5"
	"github.com/elastic/go-elasticsearch/v8"
)

type QueueElasticSearchConsumer struct{}

func (consumer *QueueElasticSearchConsumer) Consume(delivery rmq.Delivery) {
	// var message Message
	var message producer.Message

	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	esClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("\nerror when connect to elastic search")
	}

	if err := json.Unmarshal([]byte(delivery.Payload()), &message); err != nil {
		log.Printf("\nerror when parse json %v", err)

		if err := delivery.Reject(); err != nil {
			log.Printf("\nerror when rejected delivery: %v", err)
		}

		return
	}

	// perform task
	log.Printf("performing task type:%v, country_name: %v, id: %v", message.MessageType, message.Country.CountryName, message.Country.Id)

	switch message.MessageType {
	case producer.Add:
		addCountryToElasticSearch(delivery, esClient, message.Country)
	case producer.Update:
		updateCountryToElasticSearch(&delivery, esClient, message.Country)
	case producer.Delete:
		deleteCountryInElasticSearch(&delivery, esClient, message.Country)
	default:
		// reject cai nay neu khong phai la 3 truong hop, vay thi ai se xu li cai nay, no se ton tai trong queue mai mai
		log.Printf("\nelatic_search_consumer_err: 3 cases not matched")

		err := delivery.Reject()
		if err != nil {
			log.Printf("\nelatic_search_consumer_err: 3 cases not matched and rejected cause error: %v", err)
		}
	}
}

func deleteCountryInElasticSearch(delivery *rmq.Delivery, esClient *elasticsearch.TypedClient, country producer.Country) {
	res, err :=
		esClient.
			Delete(elastic_search_queue_constants.COUNTRY_INDEX, fmt.Sprintf("%s:%d", elastic_search_queue_constants.COUNTRY_INDEX, country.Id)).
			Do(context.TODO())

	if err != nil {
		log.Printf("\ndeleteCountryInElasticSearchError: error when delete document %v", err)
		if err = (*delivery).Reject(); err != nil {
			log.Printf("\ndeleteCountryInElasticSearchError: error when rejected %v", err)
		}

		return
	}

	err = (*delivery).Ack()
	if err != nil {
		log.Printf("\naddCountryToElasticSearchError: error when ack %v", err)
	}

	log.Printf("res: %v\n", res.Result.Name)
}

func updateCountryToElasticSearch(delivery *rmq.Delivery, esClient *elasticsearch.TypedClient, country producer.Country) {
	// step 1. Find id's document
	// step 2. overwrite document
	res, err := esClient.
		Index(elastic_search_queue_constants.COUNTRY_INDEX).
		Id(fmt.Sprintf("%s:%d", elastic_search_queue_constants.COUNTRY_INDEX, country.Id)).
		Request((country)).
		Do(context.TODO())

	if err != nil {
		log.Printf("\nupdateCountryToElasticSearchError: error when update document %v", err)
		if err = (*delivery).Reject(); err != nil {
			log.Printf("\nupdateCountryToElasticSearchError: error when rejected %v", err)
		}

		return
	}

	err = (*delivery).Ack()
	if err != nil {
		log.Printf("\naddCountryToElasticSearchError: error when ack %v", err)
	}

	log.Printf("res: %v\n", res.Result.Name)
}

func addCountryToElasticSearch(delivery rmq.Delivery, esClient *elasticsearch.TypedClient, country producer.Country) {
	res, err := esClient.
		Index(elastic_search_queue_constants.COUNTRY_INDEX).
		Id(fmt.Sprintf("%s:%d", elastic_search_queue_constants.COUNTRY_INDEX, country.Id)).
		Request((country)).
		Do(context.TODO())

	if err != nil {
		log.Printf("\naddCountryToElasticSearchError: error when index document %v", err)
		if err = delivery.Reject(); err != nil {
			log.Printf("\naddCountryToElasticSearchError: error when rejected %v", err)
		}

		return
	}

	err = delivery.Ack()
	if err != nil {
		log.Printf("\naddCountryToElasticSearchError: error when ack %v", err)
	}

	fmt.Printf("res: %v\n", res.Result.Name)
}
