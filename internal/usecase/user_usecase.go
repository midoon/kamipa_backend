package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/midoon/kamipa_backend/internal/domain"
	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
	"github.com/midoon/kamipa_backend/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo    domain.UserRepository
	studentRepo domain.StudentRepository
	validate    *validator.Validate
	tokenUtil   *util.TokenUtil
	redisRepo   domain.RedisRepository
}

func NewUserUsecase(validate *validator.Validate, userRepo domain.UserRepository, studentRepo domain.StudentRepository, tokenUtil *util.TokenUtil, redisRepo domain.RedisRepository) domain.UserUseCase {
	return &userUsecase{
		userRepo:    userRepo,
		studentRepo: studentRepo,
		validate:    validate,
		tokenUtil:   tokenUtil,
		redisRepo:   redisRepo,
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

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, "failed to hash password", err)
	}

	user := kamipa_entity.User{
		StudentNisn: request.StudentNisn,
		Email:       request.Email,
		Password:    string(hash),
	}

	if err = u.userRepo.Store(ctx, &user); err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, "failed to store user", err)
	}

	return nil
}

func (u *userUsecase) Login(ctx context.Context, request model.LoginUserRequest) (model.TokenDataResponse, error) {
	if err := u.validate.Struct(request); err != nil {
		return model.TokenDataResponse{}, helper.NewCustomError(http.StatusBadRequest, "validation error", err)
	}

	user, err := u.userRepo.GetByNisn(ctx, request.StudentNisn)
	if err != nil {
		return model.TokenDataResponse{}, helper.NewCustomError(http.StatusBadRequest, "failed to get user", err)
	}

	if user == (kamipa_entity.User{}) {
		return model.TokenDataResponse{}, helper.NewCustomError(http.StatusUnauthorized, "NISN or password wrong", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return model.TokenDataResponse{}, helper.NewCustomError(http.StatusUnauthorized, "NISN or password wrong", err)
	}

	redisKey := fmt.Sprintf("refresh_token:%s", user.ID)
	result, err := u.redisRepo.DeleteData(ctx, redisKey)
	if err != nil {
		return model.TokenDataResponse{}, helper.NewCustomError(http.StatusInternalServerError, "redis error :", err)
	}

	if result > 0 {
		fmt.Println("Berhadil hapus data existing refresh token from redis")
	} else {
		fmt.Println("Tidak ada data yang dihapus")
	}

	// generate token
	expAT := time.Now().Add(time.Hour * 24).UnixMilli()
	access_token, err := u.tokenUtil.CreateToken(ctx, &user, expAT)
	if err != nil {
		return model.TokenDataResponse{}, helper.NewCustomError(http.StatusInternalServerError, "failed create access token", err)
	}

	expRT := time.Now().Add(time.Hour * 24 * 30).UnixMilli()
	refresh_token, err := u.tokenUtil.CreateToken(ctx, &user, expRT)
	if err != nil {
		return model.TokenDataResponse{}, helper.NewCustomError(http.StatusInternalServerError, "failed create refresh token", err)
	}

	// store refresh token to redis db
	ttl := time.Until(time.Now().Add(time.Hour * 24 * 30))
	if err = u.redisRepo.SetDataString(ctx, redisKey, refresh_token, ttl); err != nil {
		return model.TokenDataResponse{}, helper.NewCustomError(http.StatusInternalServerError, "redis error :", err)
	}

	return model.TokenDataResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}, nil
}
