package repository

import (
	"context"

	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"gorm.io/gorm"
)

type studentRepository struct {
	simipaDB *gorm.DB
}

func NewStudentRepository(simipaDB *gorm.DB) *studentRepository {
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
