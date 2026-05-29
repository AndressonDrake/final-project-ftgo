package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(host, user, password, dbName string, port int) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=require TimeZone=Asia/Jakarta", host, user, password, dbName, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
		return
	}

	sqDB, _ := db.DB()

	err = sqDB.Ping()

	if err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
		return
	}

	return
}

func ConnectDBSlave(host, user, password, dbName, dbSchema string, port int) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=require TimeZone=Asia/Jakarta search_path = %s", host, user, password, dbName, port, dbSchema)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
		return
	}

	sqDB, _ := db.DB()

	err = sqDB.Ping()

	if err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
		return
	}

	return
}
