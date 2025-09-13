package repository

import (
	"context"

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

func (r *studentRepository) GetByNisn(ctx context.Context, nisn string) (string, error) {
	return "", nil
}
