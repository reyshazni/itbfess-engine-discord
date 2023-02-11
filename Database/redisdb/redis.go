package redisdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var redisClient *redis.Client
var ctx = context.Background()

func loadRedis() {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_DSN"))
	client := redis.NewClient(opt)
	redisClient = client
}

func GetClient() *redis.Client {
	if redisClient == nil {
		loadRedis()
	}
	return redisClient
}

func PushData(key string, values ...interface{}) error {
	err := redisClient.LPush(ctx, key, values).Err()
	return err
}

func PopData(key string) (string, error) {
	return redisClient.LPop(ctx, key).Result()
}
