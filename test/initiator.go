package test

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/usecase"
	"github.com/midoon/kamipa_backend/internal/util"
	"github.com/midoon/kamipa_backend/test/mockrepo"
	"github.com/stretchr/testify/mock"
)

type testDeps struct {
	userRepo         *mockrepo.UserRepositoryMock
	studentRepo      *mockrepo.StudentRepositoryMock
	validate         *validator.Validate
	ctx              context.Context
	redisRepo        *mockrepo.RedisRepositoryMock
	tokenUtil        *util.TokenUtil
	userUsecase      domain.UserUseCase
	dashboardRepo    *mockrepo.DashboardRepositoryMock
	dashboardUsecase domain.DashboardUsecase
}

func SetupDeps() testDeps {
	userRepo := &mockrepo.UserRepositoryMock{Mock: mock.Mock{}}
	studentRepo := &mockrepo.StudentRepositoryMock{Mock: mock.Mock{}}
	validate := validator.New()
	ctx := context.Background()
	redisRepo := &mockrepo.RedisRepositoryMock{}
	tokenUtil := util.NewTokenUtil("jwt_key", redisRepo)
	userUsecase := usecase.NewUserUsecase(validate, userRepo, studentRepo, tokenUtil, redisRepo)
	dashboardRepo := &mockrepo.DashboardRepositoryMock{Mock: mock.Mock{}}
	dashboardUsecase := usecase.NewDashboardUsecase(dashboardRepo)
	return testDeps{userRepo, studentRepo, validate, ctx, redisRepo, tokenUtil, userUsecase, dashboardRepo, dashboardUsecase}
}
