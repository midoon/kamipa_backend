package mockrepo

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type RedisRepositoryMock struct {
	mock.Mock
}

func (m *RedisRepositoryMock) GetDataString(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	val, _ := args.Get(0).(string)
	err, _ := args.Get(1).(error)
	return val, err
}

func (m *RedisRepositoryMock) SetDataString(ctx context.Context, key string, value string, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *RedisRepositoryMock) ExistData(ctx context.Context, key string) (int, error) {
	args := m.Called(ctx, key)
	count, _ := args.Get(0).(int)
	err, _ := args.Get(1).(error)
	return count, err
}

func (m *RedisRepositoryMock) DeleteData(ctx context.Context, key string) (int, error) {
	args := m.Called(ctx, key)
	del, _ := args.Get(0).(int)
	err, _ := args.Get(1).(error)
	return del, err
}
