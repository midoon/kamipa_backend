package repository

import (
	"context"
	"log"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"gorm.io/gorm"
)

type studentRepository struct {
	simipaDB *gorm.DB
}

func NewStudentRepository(simipaDB *gorm.DB) domain.StudentRepository {
	return &studentRepository{
		simipaDB: simipaDB,
	}
}

func (r *studentRepository) GetByNisn(ctx context.Context, nisn string) (simipa_entity.Student, error) {
	var student simipa_entity.Student
	err := r.simipaDB.WithContext(ctx).Where("nisn = ?", nisn).First(&student).Error
	if err != nil {
		return simipa_entity.Student{}, err
	}

	return student, nil
}

func (r *studentRepository) CountByNisn(ctx context.Context, nisn string) (int16, error) {
	var count int64
	err := r.simipaDB.WithContext(ctx).Model(&simipa_entity.Student{}).Where("nisn = ?", nisn).Count(&count).Error
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return int16(count), nil
}
