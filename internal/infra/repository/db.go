package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cestevezing/veloces/internal/core/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_CONNECTION_STRING")

	var dbAux *sql.DB
	var err error

	for range 10 {
		dbAux, err = sql.Open("mysql", dsn)
		if err == nil {
			err = dbAux.Ping()
			if err == nil {
				break
			}
		}
		log.Println("Waiting for MySQL to be ready...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MySQL successfully!")

	err = db.AutoMigrate(&model.Product{}, &model.Order{}, &model.OrderItem{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
