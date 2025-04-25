package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(host, port, dbName, username, password, sslMode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbName, username, password, sslMode,
	)

	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Close(dbConn *gorm.DB) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return
	}

	_ = sqlDB.Close()
}
