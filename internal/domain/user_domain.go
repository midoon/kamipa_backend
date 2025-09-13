package domain

import (
	"context"

	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
)

type UserRepository interface {
	Store(ctx context.Context, user *kamipa_entity.User) error
}

type UserUseCase interface {
}
