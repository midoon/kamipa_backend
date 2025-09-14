package usecase

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/midoon/kamipa_backend/internal/domain"
	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
	"github.com/midoon/kamipa_backend/internal/util"
)

type userUsecase struct {
	userRepo    domain.UserRepository
	studentRepo domain.StudentRepository
	validate    *validator.Validate
}

func NewUserRepository(validate *validator.Validate, userRepo domain.UserRepository, studentRepo domain.StudentRepository) domain.UserUseCase {
	return &userUsecase{
		userRepo:    userRepo,
		studentRepo: studentRepo,
		validate:    validate,
	}
}

func (u *userUsecase) Register(ctx context.Context, request model.RegistrationUserRequest) error {
	if err := u.validate.Struct(request); err != nil {
		return helper.NewCustomError(http.StatusBadRequest, "validation error", err)
	}

	// Check if email already exists
	count, err := u.userRepo.CountByEmail(ctx, request.Email)
	if err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, "failed to check email", err)
	}

	if count > 0 {
		return helper.NewCustomError(http.StatusConflict, "email already exists", nil)
	}

	// counte by NISN

	student, err := u.studentRepo.GetByNisn(ctx, request.StudentNisn)
	if err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, "failed to get student by NISN", err)
	}
	if student == (simipa_entity.Student{}) {
		return helper.NewCustomError(http.StatusNotFound, "student not found", nil)
	}

	hash, err := util.HashPassword(request.Password)
	if err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, "failed to hash password", err)
	}

	user := kamipa_entity.User{
		StudentNisn: request.StudentNisn,
		Email:       request.Email,
		Password:    hash,
	}

	if err = u.userRepo.Store(ctx, &user); err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, "failed to store user", err)
	}

	return nil
}
