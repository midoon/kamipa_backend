package repository

import (
	"context"
	"time"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/redis/go-redis/v9"
)

type redisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) domain.RedisRepository {
	return &redisRepository{
		redisClient: redisClient,
	}
}

func (r *redisRepository) GetDataString(ctx context.Context, key string) (string, error) {
	val, err := r.redisClient.Get(ctx, "key").Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
func (r *redisRepository) SetDataString(ctx context.Context, key string, value string, expiration time.Duration) error {

	if err := r.redisClient.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisRepository) ExistData(ctx context.Context, key string) (int, error) {
	count, err := r.redisClient.Exists(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r *redisRepository) DeleteData(ctx context.Context, key string) (int, error) {
	del, err := r.redisClient.Del(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return int(del), nil
}
