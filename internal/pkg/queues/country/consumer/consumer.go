package country_count_consumer

import (
	"encoding/json"
	"engineer-country-management/internal/pkg/mysql"
	"log"

	producer "engineer-country-management/internal/pkg/queues/country/producer"

	"github.com/adjust/rmq/v5"
)

type CountryCountConsumer struct{}

func (consumer *CountryCountConsumer) Consume(delivery rmq.Delivery) {
	db := mysql.GetClient()
	var countryCounts []producer.CountryCount
	if err := json.Unmarshal([]byte(delivery.Payload()), &countryCounts); err != nil {
		// handle json error
		log.Printf("\nerror when parse json %v", err)

		if err := delivery.Reject(); err != nil {
			// handle reject error
			log.Printf("\nerr when reject devivery :%v", err)
		}
		return
	}

	// perform task
	log.Printf("\nperforming task %v", countryCounts)
	query := producer.GetBulkUpdateQuery(&countryCounts)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("error when executed query :%v", err)
	}
	log.Printf("inserted/updated")

	if err := delivery.Ack(); err != nil {
		// handle ack error
		log.Fatalf("error orcur when ack :%v", err)
	}
}
