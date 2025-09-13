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

func (r *userRepository) CountByEmail(ctx context.Context, email string) (int16, error) {
	var count int64
	err := r.kamipaDB.WithContext(ctx).Model(&kamipa_entity.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return int16(count), nil
}
