package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/midoon/kamipa_backend/internal/configs"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/midoon/kamipa_backend/internal/model"
	"github.com/midoon/kamipa_backend/internal/usecase"
	"github.com/midoon/kamipa_backend/test/mockrepo"
	"github.com/stretchr/testify/mock"
)

func setup() (*mockrepo.UserRepositoryMock, *mockrepo.StudentRepositoryMock, *validator.Validate) {
	userRepo := &mockrepo.UserRepositoryMock{Mock: mock.Mock{}}
	studentRepo := &mockrepo.StudentRepositoryMock{Mock: mock.Mock{}}
	validate := validator.New()
	return userRepo, studentRepo, validate
}

var validate = configs.NewValidator()
var ctx = context.Background()

var userRepo = mockrepo.UserRepositoryMock{Mock: mock.Mock{}}
var studentRepo = mockrepo.StudentRepositoryMock{Mock: mock.Mock{}}
var userUsecase = usecase.NewUserUsecase(validate, &userRepo, &studentRepo)

var userRegister1 = model.RegistrationUserRequest{
	Email:       "super@gmail.com",
	Password:    "12345678",
	StudentNisn: "123",
}

func TestUserRegister(t *testing.T) {

	// mocking ngga boleh di hardcode,, dan parameternya harus sama dengan yang di usecase
	userRepo.Mock.On("Store", mock.Anything, mock.AnythingOfType("*kamipa_entity.User")).Return(nil)
	userRepo.Mock.On("CountByEmail", mock.Anything, mock.AnythingOfType("string")).Return(int16(0), nil)
	studentRepo.Mock.On("GetByNisn", mock.Anything, mock.AnythingOfType("string")).
		Return(simipa_entity.Student{}, nil)

	err := userUsecase.Register(ctx, userRegister1)
	fmt.Println(err)

}
