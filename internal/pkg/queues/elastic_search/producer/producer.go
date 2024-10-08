package queue_elastic_search_producer

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/adjust/rmq/v5"
)

type MessageType int

const (
	Add MessageType = iota
	Update
	Delete
)

type Country struct {
	Id          int64     `json:"id"`
	CountryName string    `json:"country_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Message struct {
	MessageType MessageType
	Country     Country
}

type QueueElasticSearchProducer struct {
	// conn  *rmq.Connection
	Queue *rmq.Queue
}

func (p *QueueElasticSearchProducer) AddCountryPublish(country Country) error {
	message := Message{MessageType: Add, Country: country}

	deliveryBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error when convert country to bytes")
	}

	err = (*p.Queue).PublishBytes(deliveryBytes)
	if err != nil {
		return fmt.Errorf("err when publish %v", err)
	}

	log.Printf("published: %v", message)

	return nil
}

func (p *QueueElasticSearchProducer) UpdateCountryPublish(country Country) error {
	message := Message{MessageType: Update, Country: country}

	deliveryBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error when convert country to bytes")
	}

	err = (*p.Queue).PublishBytes(deliveryBytes)
	if err != nil {
		return fmt.Errorf("err when published update message: %v", err)
	}

	log.Printf("published update message: %v", message)

	return nil
}

func (p *QueueElasticSearchProducer) DeleteCountryPublish(country Country) error {
	message := Message{MessageType: Delete, Country: country}

	deliveryBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error when convert country to bytes")
	}

	err = (*p.Queue).PublishBytes(deliveryBytes)
	if err != nil {
		return fmt.Errorf("err when published delete message: %v", err)
	}

	log.Printf("published delete message: %v", message)

	return nil
}
