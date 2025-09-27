package test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/midoon/kamipa_backend/internal/domain"
	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/midoon/kamipa_backend/internal/model"
	"github.com/midoon/kamipa_backend/internal/usecase"
	"github.com/midoon/kamipa_backend/internal/util"
	"github.com/midoon/kamipa_backend/test/mockrepo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// func setup() (*mockrepo.UserRepositoryMock, *mockrepo.StudentRepositoryMock, *validator.Validate, context.Context, *mockrepo.RedisRepositoryMock, *util.TokenUtil) {

// 	userRepo := &mockrepo.UserRepositoryMock{Mock: mock.Mock{}}
// 	studentRepo := &mockrepo.StudentRepositoryMock{Mock: mock.Mock{}}
// 	validate := validator.New()
// 	ctx := context.Background()
// 	redisRepo := &mockrepo.RedisRepositoryMock{}

// 	tokenUtil := util.NewTokenUtil("rahasia", redisRepo)

// 	return userRepo, studentRepo, validate, ctx, redisRepo, tokenUtil
// }

type testDeps struct {
	userRepo    *mockrepo.UserRepositoryMock
	studentRepo *mockrepo.StudentRepositoryMock
	validate    *validator.Validate
	ctx         context.Context
	redisRepo   *mockrepo.RedisRepositoryMock
	tokenUtil   *util.TokenUtil
	userUsecase domain.UserUseCase
}

func setupDeps() testDeps {
	userRepo := &mockrepo.UserRepositoryMock{Mock: mock.Mock{}}
	studentRepo := &mockrepo.StudentRepositoryMock{Mock: mock.Mock{}}
	validate := validator.New()
	ctx := context.Background()
	redisRepo := &mockrepo.RedisRepositoryMock{}
	tokenUtil := util.NewTokenUtil("jwt_key", redisRepo)
	userUsecase := usecase.NewUserUsecase(validate, userRepo, studentRepo, tokenUtil, redisRepo)
	return testDeps{userRepo, studentRepo, validate, ctx, redisRepo, tokenUtil, userUsecase}
}

func TestUserRegister(t *testing.T) {

	userRegister1 := model.RegistrationUserRequest{
		Email:       "user1@gmail.com",
		Password:    "12345678",
		StudentNisn: "1312",
	}

	t.Run("success register", func(t *testing.T) {
		deps := setupDeps()

		// mocking ngga boleh di hardcode,, dan parameternya harus sama dengan yang di usecase
		deps.userRepo.Mock.On("Store", mock.Anything, mock.AnythingOfType("*kamipa_entity.User")).Return(nil)
		deps.userRepo.Mock.On("CountByEmail", mock.Anything, mock.AnythingOfType("string")).Return(int16(0), nil)
		deps.studentRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).
			Return(simipa_entity.Student{
				ID:      1,
				GroupId: 1,
				Name:    "student-name",
				Nisn:    "1312",
				Gender:  "laki-laki",
			}, nil)

		err := deps.userUsecase.Register(deps.ctx, userRegister1)
		assert.NoError(t, err)
		deps.userRepo.Mock.AssertExpectations(t)
		deps.studentRepo.Mock.AssertExpectations(t)
	})

	t.Run("error validation", func(t *testing.T) {
		deps := setupDeps()

		invalidReq := model.RegistrationUserRequest{} // kosong semua -> invalid

		err := deps.userUsecase.Register(deps.ctx, invalidReq)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "validation error")
	})

	t.Run("duplicate user", func(t *testing.T) {
		deps := setupDeps()

		deps.userRepo.Mock.On("Store", mock.Anything, mock.AnythingOfType("*kamipa_entity.User")).Return(nil)
		deps.userRepo.Mock.On("CountByEmail", mock.Anything, mock.AnythingOfType("string")).Return(int16(1), nil)
		deps.studentRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).
			Return(simipa_entity.Student{
				ID:      1,
				GroupId: 1,
				Name:    "student-name",
				Nisn:    "1312",
				Gender:  "laki-laki",
			}, nil)

		err := deps.userUsecase.Register(deps.ctx, userRegister1)
		assert.Contains(t, err.Error(), "email already exists")
	})

	t.Run("wrong student nisn", func(t *testing.T) {
		deps := setupDeps()

		deps.userRepo.Mock.On("Store", mock.Anything, mock.AnythingOfType("*kamipa_entity.User")).Return(nil)
		deps.userRepo.Mock.On("CountByEmail", mock.Anything, mock.AnythingOfType("string")).Return(int16(0), nil)
		deps.studentRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).
			Return(simipa_entity.Student{}, errors.New("record not found"))

		err := deps.userUsecase.Register(deps.ctx, userRegister1)
		assert.Contains(t, err.Error(), "failed to get student by NISN")
	})

}

func TestUserLogin(t *testing.T) {

	userLogin1 := model.LoginUserRequest{

		Password:    "12345678",
		StudentNisn: "1312",
	}

	t.Run("success login", func(t *testing.T) {
		deps := setupDeps()

		// hash password
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

		deps.userRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).Return(kamipa_entity.User{
			ID:          "id-1",
			StudentNisn: "1312",
			Email:       "user1@gmail.com",
			Password:    string(hashedPassword),
		}, nil)

		// mock redis
		deps.redisRepo.Mock.On("DeleteData", mock.Anything, mock.AnythingOfType("string")).Return(1, nil)
		deps.redisRepo.Mock.On("SetDataString",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("time.Duration"),
		).Return(nil)

		tokenData, err := deps.userUsecase.Login(deps.ctx, userLogin1)
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenData.AccessToken)
		assert.NotEmpty(t, tokenData.RefreshToken)
		deps.userRepo.Mock.AssertExpectations(t)
		deps.redisRepo.Mock.AssertExpectations(t)
	})

	t.Run("validation error", func(t *testing.T) {
		deps := setupDeps()
		// request kosong supaya kena validator
		invalidReq := model.LoginUserRequest{}
		tokenData, err := deps.userUsecase.Login(deps.ctx, invalidReq)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "validation error")
		assert.Empty(t, tokenData.AccessToken)
	})

	t.Run("failed get user (repo error)", func(t *testing.T) {
		deps := setupDeps()
		// mock GetByNisn return error
		deps.userRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).Return(kamipa_entity.User{}, errors.New("db down"))

		req := model.LoginUserRequest{
			StudentNisn: "1312",
			Password:    "12345678",
		}

		tokenData, err := deps.userUsecase.Login(deps.ctx, req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get user")
		assert.Empty(t, tokenData.AccessToken)
	})

	t.Run("wrong NISN (user kosong)", func(t *testing.T) {
		deps := setupDeps()
		// mock GetByNisn return user kosong tanpa error
		deps.userRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).Return(kamipa_entity.User{}, nil)

		req := model.LoginUserRequest{
			StudentNisn: "1312",
			Password:    "12345678",
		}

		tokenData, err := deps.userUsecase.Login(deps.ctx, req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "NISN or password wrong")
		assert.Empty(t, tokenData.AccessToken)
	})

	t.Run("wrong password", func(t *testing.T) {
		deps := setupDeps()

		// hash password beda
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("passwordlain"), bcrypt.DefaultCost)

		deps.userRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).Return(kamipa_entity.User{
			ID:          "id-1",
			StudentNisn: "1312",
			Email:       "user1@gmail.com",
			Password:    string(hashedPassword),
		}, nil)

		req := model.LoginUserRequest{
			StudentNisn: "1312",
			Password:    "12345678", // password salah
		}

		tokenData, err := deps.userUsecase.Login(deps.ctx, req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "NISN or password wrong")
		assert.Empty(t, tokenData.AccessToken)
	})

	t.Run("redis error on delete", func(t *testing.T) {
		deps := setupDeps()

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

		deps.userRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).Return(kamipa_entity.User{
			ID:          "id-1",
			StudentNisn: "1312",
			Email:       "user1@gmail.com",
			Password:    string(hashedPassword),
		}, nil)

		// redis delete error
		deps.redisRepo.Mock.On("DeleteData", mock.Anything, mock.AnythingOfType("string")).Return(0, errors.New("redis down"))

		req := model.LoginUserRequest{
			StudentNisn: "1312",
			Password:    "12345678",
		}

		tokenData, err := deps.userUsecase.Login(deps.ctx, req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "redis error")
		assert.Empty(t, tokenData.AccessToken)
	})

	t.Run("redis error on set refresh token", func(t *testing.T) {
		deps := setupDeps()

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

		deps.userRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).Return(kamipa_entity.User{
			ID:          "id-1",
			StudentNisn: "1312",
			Email:       "user1@gmail.com",
			Password:    string(hashedPassword),
		}, nil)

		deps.redisRepo.Mock.On("DeleteData", mock.Anything, mock.AnythingOfType("string")).Return(1, nil)
		// SetDataString error
		deps.redisRepo.Mock.On("SetDataString",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("time.Duration"),
		).Return(errors.New("redis down"))

		req := model.LoginUserRequest{
			StudentNisn: "1312",
			Password:    "12345678",
		}

		tokenData, err := deps.userUsecase.Login(deps.ctx, req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "redis error")
		assert.Empty(t, tokenData.AccessToken)
	})
}
