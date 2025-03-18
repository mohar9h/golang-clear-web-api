package main

import (
	"github.com/mohar9h/golang-clear-web-api/api"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/data/cache"
	"github.com/mohar9h/golang-clear-web-api/data/db"
	"github.com/mohar9h/golang-clear-web-api/logging"
)

// @securityDefinitions.api_key Bearer
// @in header
// @name Authorization
func main() {
	config := config.GetConfig()

	logger := logging.NewLogger(config)

	err := cache.InitRedis(config)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	err = db.InitDatabase(config)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDB()
	api.InitServer(config)
}
