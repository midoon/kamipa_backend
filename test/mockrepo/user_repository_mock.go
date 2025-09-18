package mockrepo

import (
	"context"

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
	return arguments.Error(0)
}
func (r *UserRepositoryMock) CountByEmail(ctx context.Context, email string) (int16, error) {
	// 	r.Mock.Called mengembalikan nilai return yang sudah lo set di .On().Return(...)
	// Jadi arguments[0] = return pertama, arguments[1] = return kedua.
	// Bukan argumen yang dikirim ke fungsi.
	arguments := r.Mock.Called(ctx, email)

	count, _ := arguments.Get(0).(int16)
	err, _ := arguments.Get(1).(error)

	return count, err
}

func (r *UserRepositoryMock) GetByNisn(ctx context.Context, nisn string) (kamipa_entity.User, error) {
	args := r.Mock.Called(ctx, nisn)

	user, _ := args.Get(0).(kamipa_entity.User)
	err, _ := args.Get(1).(error)

	return user, err
}
