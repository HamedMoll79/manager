package redisqueue

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var c *redis.Client

func New(config Config, name string) *redis.Client {
	c = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return c
}

func FetchMessage(queueName string) {
	data, err := zc.LPop(context.Background(), queueName).Result()
}
