package test

import (
	"errors"
	"testing"

	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/midoon/kamipa_backend/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

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
		deps.studentRepo.Mock.On("CountByNisn", mock.Anything, mock.AnythingOfType("string")).Return(int16(1), nil)
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
		deps.studentRepo.Mock.On("CountByNisn", mock.Anything, mock.AnythingOfType("string")).Return(int16(0), nil)
		deps.studentRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).
			Return(simipa_entity.Student{}, errors.New("record not found"))

		err := deps.userUsecase.Register(deps.ctx, userRegister1)
		assert.Contains(t, err.Error(), "student not found")
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
		deps.userRepo.Mock.On("CountByNisn", mock.Anything, mock.AnythingOfType("string")).Return(int16(1), nil)

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
		deps.userRepo.Mock.On("CountByNisn", mock.Anything, mock.AnythingOfType("string")).Return(int16(1), nil)
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
		deps.userRepo.Mock.On("CountByNisn", mock.Anything, mock.AnythingOfType("string")).Return(int16(0), nil)
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

		deps.userRepo.Mock.On("CountByNisn", mock.Anything, mock.AnythingOfType("string")).Return(int16(1), nil)
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

		deps.userRepo.Mock.On("CountByNisn", mock.Anything, mock.AnythingOfType("string")).Return(int16(1), nil)
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

		deps.userRepo.Mock.On("CountByNisn", mock.Anything, mock.AnythingOfType("string")).Return(int16(1), nil)
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

func TestRefreshToken(t *testing.T) {
	password := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	tempDeps := setupDeps()
	tempDeps.userRepo.Mock.On("CountByNisn", mock.Anything, mock.AnythingOfType("string")).Return(int16(1), nil)
	tempDeps.userRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).
		Return(kamipa_entity.User{
			ID:          "id-1",
			StudentNisn: "1312",
			Email:       "user1@gmail.com",
			Password:    string(hashedPassword),
		}, nil)
	tempDeps.redisRepo.Mock.On("DeleteData", mock.Anything, mock.AnythingOfType("string")).Return(1, nil)
	tempDeps.redisRepo.Mock.On("SetDataString", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	tokenData, err := tempDeps.userUsecase.Login(tempDeps.ctx, model.LoginUserRequest{
		Password:    password,
		StudentNisn: "1312",
	})
	require.NoError(t, err)

	t.Run("success refresh token", func(t *testing.T) {
		deps := setupDeps()
		deps.userRepo.Mock.On("GetById", mock.Anything, "id-1").Return(kamipa_entity.User{
			ID:          "id-1",
			StudentNisn: "1312",
			Email:       "user1@gmail.com",
			Password:    string(hashedPassword),
		}, nil)
		deps.redisRepo.Mock.On("ExistData", mock.Anything, mock.AnythingOfType("string")).Return(1, nil)

		request := model.RefreshTokenRequest{RefreshToken: tokenData.RefreshToken}
		newTokenData, err := deps.userUsecase.RefreshToken(deps.ctx, request)

		assert.NoError(t, err)
		assert.NotEmpty(t, newTokenData.AccessToken)
		assert.NotEmpty(t, newTokenData.RefreshToken)

		deps.userRepo.Mock.AssertExpectations(t)
		deps.redisRepo.Mock.AssertExpectations(t)
	})

	t.Run("validation error", func(t *testing.T) {
		deps := setupDeps()

		request := model.RefreshTokenRequest{
			RefreshToken: "", // invalid, kosong
		}

		_, err := deps.userUsecase.RefreshToken(deps.ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "validation error")
	})

	t.Run("invalid refresh token", func(t *testing.T) {
		deps := setupDeps()

		request := model.RefreshTokenRequest{
			RefreshToken: "token-salah",
		}

		_, err := deps.userUsecase.RefreshToken(deps.ctx, request)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Unauthorized")
	})

	t.Run("redis error", func(t *testing.T) {
		deps := setupDeps()
		deps.userRepo.Mock.On("GetById", mock.Anything, "id-1").Return(kamipa_entity.User{
			ID:          "id-1",
			StudentNisn: "1312",
			Email:       "user1@gmail.com",
			Password:    string(hashedPassword),
		}, nil)
		deps.redisRepo.Mock.On("ExistData", mock.Anything, mock.AnythingOfType("string")).Return(1, errors.New("redis down"))

		request := model.RefreshTokenRequest{RefreshToken: tokenData.RefreshToken}
		_, err := deps.userUsecase.RefreshToken(deps.ctx, request)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "redis error")
	})

	t.Run("refresh token not found in redis", func(t *testing.T) {
		deps := setupDeps()
		deps.userRepo.Mock.On("GetById", mock.Anything, "id-1").Return(kamipa_entity.User{
			ID:          "id-1",
			StudentNisn: "1312",
			Email:       "user1@gmail.com",
			Password:    string(hashedPassword),
		}, nil)
		deps.redisRepo.Mock.On("ExistData", mock.Anything, mock.AnythingOfType("string")).Return(0, nil)

		request := model.RefreshTokenRequest{RefreshToken: tokenData.RefreshToken}
		_, err := deps.userUsecase.RefreshToken(deps.ctx, request)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Unauthorized")
	})

	t.Run("user not found in database", func(t *testing.T) {
		deps := setupDeps()
		deps.userRepo.Mock.On("GetById", mock.Anything, "id-1").Return(kamipa_entity.User{}, errors.New("record not found"))
		deps.redisRepo.Mock.On("ExistData", mock.Anything, mock.AnythingOfType("string")).Return(1, nil)

		request := model.RefreshTokenRequest{RefreshToken: tokenData.RefreshToken}
		_, err := deps.userUsecase.RefreshToken(deps.ctx, request)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get user")
	})

}

func TestLogout(t *testing.T) {
	t.Run("success logout", func(t *testing.T) {
		deps := setupDeps()
		deps.redisRepo.Mock.On("DeleteData", mock.Anything, mock.AnythingOfType("string")).Return(1, nil)

		err := deps.userUsecase.Logout(deps.ctx, "id-1")
		assert.NoError(t, err)
		deps.redisRepo.Mock.AssertExpectations(t)
	})

	t.Run("redis error", func(t *testing.T) {
		deps := setupDeps()
		deps.redisRepo.Mock.On("DeleteData", mock.Anything, mock.AnythingOfType("string")).Return(0, errors.New("redis down"))

		err := deps.userUsecase.Logout(deps.ctx, "id-1")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "redis error")

	})

	t.Run("no session", func(t *testing.T) {
		deps := setupDeps()
		deps.redisRepo.Mock.On("DeleteData", mock.Anything, mock.AnythingOfType("string")).Return(0, nil)

		err := deps.userUsecase.Logout(deps.ctx, "id-1")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no active session found")

	})
}
