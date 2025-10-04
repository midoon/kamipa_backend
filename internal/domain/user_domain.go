package domain

import (
	"context"

	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/model"
)

type UserRepository interface {
	Store(ctx context.Context, user *kamipa_entity.User) error
	CountByEmail(ctx context.Context, email string) (int16, error)
	CountByNisn(ctx context.Context, nisn string) (int16, error)
	GetByNisn(ctx context.Context, nisn string) (kamipa_entity.User, error)
	GetById(ctx context.Context, id string) (kamipa_entity.User, error)
}

type UserUseCase interface {
	Register(ctx context.Context, request model.RegistrationUserRequest) error
	Login(ctx context.Context, request model.LoginUserRequest) (model.TokenDataResponse, error)
	Logout(ctx context.Context, userId string) error
	RefreshToken(ctx context.Context, request model.RefreshTokenRequest) (model.TokenDataResponse, error)
}
