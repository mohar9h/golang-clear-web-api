package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/mohar9h/golang-clear-web-api/config"
)

var redisClient *redis.Client

func InitRedis(config *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password:           config.Redis.Password,
		DB:                 0,
		DialTimeout:        config.Redis.DialTimeout * time.Second,
		ReadTimeout:        config.Redis.ReadTimeout * time.Second,
		WriteTimeout:       config.Redis.WriteTimeout * time.Second,
		PoolSize:           config.Redis.PoolSize,
		PoolTimeout:        config.Redis.PoolTimeout,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: config.Redis.IdleCheckFrequency * time.Second,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}
