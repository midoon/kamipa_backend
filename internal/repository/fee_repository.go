package repository

import (
	"context"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"gorm.io/gorm"
)

type feeRepository struct {
	simipaDB *gorm.DB
}

func NewFeeRepository(simipaDB *gorm.DB) domain.FeeRepository {
	return &feeRepository{
		simipaDB: simipaDB,
	}
}

func (r *feeRepository) GetByStudentId(ctx context.Context, studentId int64) ([]simipa_entity.Fee, error) {
	var fees []simipa_entity.Fee

	err := r.simipaDB.WithContext(ctx).Preload("PaymentType").Where("student_id = ?", studentId).Order("due_date ASC").Find(&fees).Error

	if err != nil {
		return []simipa_entity.Fee{}, err
	}

	return fees, nil
}

func (r *feeRepository) GetByFeeId(ctx context.Context, feeId int64) (simipa_entity.Fee, error) {
	var fee simipa_entity.Fee
	err := r.simipaDB.WithContext(ctx).Preload("PaymentType").Where("id = ?", feeId).First(&fee).Error

	if err != nil {
		return simipa_entity.Fee{}, err
	}
	return fee, nil
}
