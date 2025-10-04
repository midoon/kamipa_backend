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

func (r *userRepository) CountByNisn(ctx context.Context, nisn string) (int16, error) {
	var count int64
	err := r.kamipaDB.WithContext(ctx).Model(&kamipa_entity.User{}).Where("student_nisn = ?", nisn).Count(&count).Error
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return int16(count), nil
}

func (r *userRepository) GetByNisn(ctx context.Context, nisn string) (kamipa_entity.User, error) {
	var user kamipa_entity.User

	err := r.kamipaDB.WithContext(ctx).Where("student_nisn = ?", nisn).First(&user).Error
	if err != nil {
		return kamipa_entity.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetById(ctx context.Context, id string) (kamipa_entity.User, error) {
	var user kamipa_entity.User

	err := r.kamipaDB.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return kamipa_entity.User{}, err
	}

	return user, nil
}
