package country_count_producer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/adjust/rmq/v5"
	"github.com/redis/go-redis/v9"
)

type CountryCount struct {
	Key   string
	Value int64
}

type CountryCountProducer struct {
	RedisClient *redis.Client
}

func GetBulkUpdateQuery(countryCounts *[]CountryCount) string {
	// Tạo câu lệnh SQL cho bulk insert/update
	var sqlStr string

	for _, cc := range *countryCounts {
		sqlStr += fmt.Sprintf("('%v', %v),", cc.Key, cc.Value)
	}

	// Loại bỏ dấu phẩy cuối cùng
	sqlStr = sqlStr[:len(sqlStr)-1]

	// Câu lệnh SQL
	query := fmt.Sprintf("INSERT INTO country_counts (`key`, value) VALUES %s ON DUPLICATE KEY UPDATE value = VALUES(value)", sqlStr)

	// fmt.Println(query)
	return query
}

func (p *CountryCountProducer) GetCountryCount() []CountryCount {
	var countryCounts []CountryCount

	// Sử dụng SCAN để tìm tất cả các keys phù hợp với pattern
	pattern := "count:country:*"
	var cursor uint64
	for {
		var keys []string
		var err error

		// Thực hiện SCAN
		keys, cursor, err = p.RedisClient.Scan(context.Background(), cursor, pattern, 0).Result()
		if err != nil {
			log.Fatalf("could not scan keys: %v", err)
		}

		// Xử lý các keys tìm được
		for _, key := range keys {
			var countCount CountryCount
			fmt.Println("Found key:", key)

			// Lấy giá trị của key
			value, err := p.RedisClient.Get(context.Background(), key).Result()
			if err != nil {
				log.Printf("could not get value for key %s: %v", key, err)
				continue
			}
			fmt.Printf("Value for key %s: %s\n", key, value)

			valueToInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				log.Printf("error when convert string redis to int64 value %v", err)
				continue
			}

			countCount.Key = key
			countCount.Value = valueToInt

			countryCounts = append(countryCounts, countCount)
		}

		// Nếu cursor là 0, đã quét hết các key
		if cursor == 0 {
			break
		}
	}

	fmt.Println("\n=======================================")

	return countryCounts
}

func (p *CountryCountProducer) CountryCountPublish(conn *rmq.Connection, queue *rmq.Queue) error {
	countryCounts := p.GetCountryCount()

	if len(countryCounts) == 0 {
		return fmt.Errorf("country counts is empty")
	}

	deliveryBytes, err := json.Marshal(countryCounts)
	if err != nil {
		return fmt.Errorf("error when convert struct delivery to bytes %v", err)
	}

	err = (*queue).PublishBytes(deliveryBytes)
	if err != nil {
		return fmt.Errorf("error when publish %v", err)
	}

	log.Printf("Published: %v", countryCounts)

	return nil
}
