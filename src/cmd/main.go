package main

import (
	"github.com/mohar9h/golang-clear-web-api/api"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/data/cache"
)

func main() {
	config := config.GetConfig()
	cache.InitRedis(config)
	defer cache.CloseRedis()
	api.InitServer(config)
}
