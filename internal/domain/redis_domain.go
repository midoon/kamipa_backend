package domain

import (
	"context"
	"time"
)

type RedisRepository interface {
	SetDataString(ctx context.Context, key string, value string, expiration time.Duration) error
	GetDataString(ctx context.Context, key string) (string, error)
	ExistData(ctx context.Context, key string) (int, error)
	DeleteData(ctx context.Context, key string) (int, error)
}
