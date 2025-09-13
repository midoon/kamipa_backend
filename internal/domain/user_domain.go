package domain

import (
	"context"

	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/model"
)

type UserRepository interface {
	Store(ctx context.Context, user *kamipa_entity.User) error
	CountByEmail(ctx context.Context, email string) (int16, error)
}

type UserUseCase interface {
	Register(ctx context.Context, request model.RegistrationUserRequest) error
}
