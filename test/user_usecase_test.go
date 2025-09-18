package test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/midoon/kamipa_backend/internal/configs"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/midoon/kamipa_backend/internal/model"
	"github.com/midoon/kamipa_backend/internal/usecase"
	"github.com/midoon/kamipa_backend/internal/util"
	"github.com/midoon/kamipa_backend/test/mockrepo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup() (*mockrepo.UserRepositoryMock, *mockrepo.StudentRepositoryMock, *validator.Validate, context.Context, *mockrepo.RedisRepositoryMock, *util.TokenUtil) {
	config := configs.GetConfig()
	userRepo := &mockrepo.UserRepositoryMock{Mock: mock.Mock{}}
	studentRepo := &mockrepo.StudentRepositoryMock{Mock: mock.Mock{}}
	validate := validator.New()
	ctx := context.Background()
	redisRepo := &mockrepo.RedisRepositoryMock{}
	tokenUtil := util.NewTokenUtil(config.JWT.Key, redisRepo)
	return userRepo, studentRepo, validate, ctx, redisRepo, tokenUtil
}

func TestUserRegister(t *testing.T) {

	userRegister1 := model.RegistrationUserRequest{
		Email:       "user1@gmail.com",
		Password:    "12345678",
		StudentNisn: "1312",
	}

	t.Run("success register", func(t *testing.T) {
		userRepo, studentRepo, validate, ctx, redisRepo, tokenUtil := setup()
		userUsecase := usecase.NewUserUsecase(validate, userRepo, studentRepo, tokenUtil, redisRepo)
		// mocking ngga boleh di hardcode,, dan parameternya harus sama dengan yang di usecase
		userRepo.Mock.On("Store", mock.Anything, mock.AnythingOfType("*kamipa_entity.User")).Return(nil)
		userRepo.Mock.On("CountByEmail", mock.Anything, mock.AnythingOfType("string")).Return(int16(0), nil)
		studentRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).
			Return(simipa_entity.Student{
				ID:      1,
				GroupId: 1,
				Name:    "student-name",
				Nisn:    "1312",
				Gender:  "laki-laki",
			}, nil)

		err := userUsecase.Register(ctx, userRegister1)
		assert.NoError(t, err)
		userRepo.Mock.AssertExpectations(t)
		studentRepo.Mock.AssertExpectations(t)
	})

	t.Run("error validation", func(t *testing.T) {
		userRepo, studentRepo, validate, ctx, redisRepo, tokenUtil := setup()
		userUsecase := usecase.NewUserUsecase(validate, userRepo, studentRepo, tokenUtil, redisRepo)

		invalidReq := model.RegistrationUserRequest{} // kosong semua -> invalid

		err := userUsecase.Register(ctx, invalidReq)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "validation error")
	})

	t.Run("duplicate user", func(t *testing.T) {
		userRepo, studentRepo, validate, ctx, redisRepo, tokenUtil := setup()
		userUsecase := usecase.NewUserUsecase(validate, userRepo, studentRepo, tokenUtil, redisRepo)

		userRepo.Mock.On("Store", mock.Anything, mock.AnythingOfType("*kamipa_entity.User")).Return(nil)
		userRepo.Mock.On("CountByEmail", mock.Anything, mock.AnythingOfType("string")).Return(int16(1), nil)
		studentRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).
			Return(simipa_entity.Student{
				ID:      1,
				GroupId: 1,
				Name:    "student-name",
				Nisn:    "1312",
				Gender:  "laki-laki",
			}, nil)

		err := userUsecase.Register(ctx, userRegister1)
		assert.Contains(t, err.Error(), "email already exists")
	})

	t.Run("wrong student nisn", func(t *testing.T) {
		userRepo, studentRepo, validate, ctx, redisRepo, tokenUtil := setup()
		userUsecase := usecase.NewUserUsecase(validate, userRepo, studentRepo, tokenUtil, redisRepo)

		userRepo.Mock.On("Store", mock.Anything, mock.AnythingOfType("*kamipa_entity.User")).Return(nil)
		userRepo.Mock.On("CountByEmail", mock.Anything, mock.AnythingOfType("string")).Return(int16(0), nil)
		studentRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).
			Return(simipa_entity.Student{}, errors.New("record not found"))

		err := userUsecase.Register(ctx, userRegister1)
		assert.Contains(t, err.Error(), "failed to get student by NISN")
	})

}
