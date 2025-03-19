package main

import (
	"github.com/mohar9h/golang-clear-web-api/api"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/data/cache"
	"github.com/mohar9h/golang-clear-web-api/data/db"
	"github.com/mohar9h/golang-clear-web-api/data/db/migrations"
	"github.com/mohar9h/golang-clear-web-api/logging"
)

// @securityDefinitions.api_key Bearer
// @in header
// @name Authorization
func main() {
	getConfig := config.GetConfig()

	logger := logging.NewLogger(getConfig)

	err := cache.InitRedis(getConfig)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	err = db.InitDatabase(getConfig)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDB()

	migrations.Up()

	api.InitServer(getConfig)
}
