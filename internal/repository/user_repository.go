package repository

import (
	"context"

	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"gorm.io/gorm"
)

type userRepository struct {
	kamipaDB *gorm.DB
}

func NewUserRepository(kamipaDB *gorm.DB) *userRepository {
	return &userRepository{
		kamipaDB: kamipaDB,
	}
}

func (r *userRepository) Store(ctx context.Context, user *kamipa_entity.User) error {
	return nil
}
