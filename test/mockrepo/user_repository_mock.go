package mockrepo

import (
	"context"
	"errors"

	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (r *UserRepositoryMock) Store(ctx context.Context, user *kamipa_entity.User) error {
	// 	r.Mock.Called mengembalikan nilai return yang sudah lo set di .On().Return(...)
	// Jadi arguments[0] = return pertama, arguments[1] = return kedua.
	// Bukan argumen yang dikirim ke fungsi.
	arguments := r.Mock.Called(ctx, user)
	// jika ada error
	if arguments.Get(0) != nil {
		return errors.New("DB ERROR")
	}

	return nil
}
func (r *UserRepositoryMock) CountByEmail(ctx context.Context, email string) (int16, error) {
	// 	r.Mock.Called mengembalikan nilai return yang sudah lo set di .On().Return(...)
	// Jadi arguments[0] = return pertama, arguments[1] = return kedua.
	// Bukan argumen yang dikirim ke fungsi.
	arguments := r.Mock.Called(ctx, email)

	// jika return value ada error
	if arguments.Get(1) != nil {
		return 0, errors.New("DB ERROR")
	}

	if arguments.Get(0) != 0 {
		return arguments.Get(0).(int16), nil
	}

	return 0, nil
}
