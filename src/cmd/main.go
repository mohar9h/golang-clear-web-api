package main

import (
	"github.com/mohar9h/golang-clear-web-api/api"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/domains/cache"
	"github.com/mohar9h/golang-clear-web-api/domains/db/migrations"
	"github.com/mohar9h/golang-clear-web-api/infrastructure/persistence/database"
	logging2 "github.com/mohar9h/golang-clear-web-api/pkg/logging"
)

// @securityDefinitions.api_key Bearer
// @in header
// @name Authorization
func main() {
	getConfig := config.GetConfig()

	logger := logging2.NewLogger(getConfig)

	err := cache.InitRedis(getConfig)
	if err != nil {
		logger.Fatal(logging2.Redis, logging2.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	err = database.InitDatabase(getConfig)
	if err != nil {
		logger.Fatal(logging2.Postgres, logging2.Startup, err.Error(), nil)
	}
	defer database.CloseDB()

	migrations.Up()

	api.InitServer(getConfig)
}
