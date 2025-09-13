package repository

import (
	"context"
	"log"

	"github.com/midoon/kamipa_backend/internal/domain"
	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"gorm.io/gorm"
)

type userRepository struct {
	kamipaDB *gorm.DB
}

func NewUserRepository(kamipaDB *gorm.DB) domain.UserRepository {
	return &userRepository{
		kamipaDB: kamipaDB,
	}
}

func (r *userRepository) Store(ctx context.Context, user *kamipa_entity.User) error {
	err := r.kamipaDB.WithContext(ctx).Create(user).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
