package redis

import (
	"context"
	"time"

	"github.com/vins7/super-indo/config"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	conf := config.GetConfig()
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Database.Redis.Host + ":" + conf.Database.Redis.Port,
		Password: "",
		DB:       0,
	})
}

func GetCache(ctx context.Context, key string) (string, error) {

	val, err := client.Get(key).Result()
	if err == redis.Nil {
		return "data not found", err
	} else if err != nil {
		return "", err
	} else {
		return val, nil
	}
}

func SetCache(ctx context.Context, key string, val string, ttl time.Duration) error {
	return client.Set(key, val, ttl).Err()
}

func DelCache(ctx context.Context, keys ...string) error {
	return client.Del(keys...).Err()
}
