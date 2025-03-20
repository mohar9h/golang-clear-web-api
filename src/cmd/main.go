package main

import (
	"github.com/mohar9h/golang-clear-web-api/api"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/data/cache"
	"github.com/mohar9h/golang-clear-web-api/data/db"
	"github.com/mohar9h/golang-clear-web-api/data/db/migrations"
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

	err = db.InitDatabase(getConfig)
	if err != nil {
		logger.Fatal(logging2.Postgres, logging2.Startup, err.Error(), nil)
	}
	defer db.CloseDB()

	migrations.Up()

	api.InitServer(getConfig)
}
