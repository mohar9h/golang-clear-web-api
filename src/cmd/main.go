package main

import (
	"log"

	"github.com/mohar9h/golang-clear-web-api/api"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/data/cache"
	"github.com/mohar9h/golang-clear-web-api/data/db"
)

func main() {
	config := config.GetConfig()
	err := cache.InitRedis(config)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()

	err = db.InitDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()
	api.InitServer(config)
}
