package configs

import "github.com/redis/go-redis/v9"

func GetRedisClient(addr string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
	})
}
