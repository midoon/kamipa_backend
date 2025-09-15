package mockrepo

import (
	"context"
	"net/http"

	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/helper"
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
	if arguments.Get(0) == nil {
		return helper.NewCustomError(http.StatusInternalServerError, "failed to store user", nil)
	}

	return nil
}
func (r *UserRepositoryMock) CountByEmail(ctx context.Context, email string) (int16, error) {
	// 	r.Mock.Called mengembalikan nilai return yang sudah lo set di .On().Return(...)
	// Jadi arguments[0] = return pertama, arguments[1] = return kedua.
	// Bukan argumen yang dikirim ke fungsi.
	arguments := r.Mock.Called(ctx, email)
	if arguments.Get(0) == nil {
		return 0, helper.NewCustomError(http.StatusInternalServerError, "failed to store user", nil)
	}

	return 0, nil
}
