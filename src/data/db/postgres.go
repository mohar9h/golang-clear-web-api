package db

import (
	"fmt"
	"log"

	"github.com/mohar9h/golang-clear-web-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDatabase(config *config.Config) error {
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.DbName, config.Postgres.Password)

	dbClient, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(config.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(config.Postgres.MaxIdleConns)
	sqlDb.SetConnMaxLifetime(config.Postgres.ConnMaxLifetime)

	log.Println("Database connection established")
	return nil
}

func GetDBClient() *gorm.DB {
	return dbClient
}

func CloseDB() {
	sqlDb, _ := dbClient.DB()
	sqlDb.Close()
}
