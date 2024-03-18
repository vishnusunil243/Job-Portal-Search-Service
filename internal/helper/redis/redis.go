package redis

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
}
func Add(userId, value string) (int64, error) {
	key := fmt.Sprintf("%s:%s", userId, value)
	result, err := redisClient.Incr(key).Result()
	if err != nil {
		return 0, err
	}

	if result == 1 {
		redisClient.Expire(key, 48*time.Hour)
	}
	if result == 5 {
		redisClient.Set(key, "1", 48*time.Hour)
	}

	return result, nil
}
func Get(userId, value string) (int64, error) {
	key := fmt.Sprintf("%s:%s", userId, value)
	num, err := redisClient.Get(key).Result()
	if err == redis.Nil {
		return 0, nil
	}
	res, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}
