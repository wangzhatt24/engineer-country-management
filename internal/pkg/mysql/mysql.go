package mysql

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once     sync.Once
	instance *sql.DB
	err      error
)

func GetClient() *sql.DB {
	once.Do(func() {
		dsn := "tyler:abc@123@tcp(127.0.0.1:3306)/engineer-country?parseTime=true"
		instance, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal("failed to connect to database: %w", err)
		}

		// Kiểm tra kết nối
		if err = instance.Ping(); err != nil {
			log.Fatal("failed to connect to database: %w", err)
		}
	})

	return instance
}

func CloseDB() error {
	if instance != nil {
		return instance.Close()
	}
	return nil
}
