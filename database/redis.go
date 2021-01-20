package database

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type IRedisConnection interface {
	CreateConnection() *redis.Client
}

type RedisConnection struct{}

func (c *RedisConnection) CreateConnection() *redis.Client {
	opt, err := redis.ParseURL(fmt.Sprintf("%v", viper.Get("redis")))
	if err != nil {
		log.Fatal(err.Error())
	}

	rdb := redis.NewClient(opt)

	return rdb
}
