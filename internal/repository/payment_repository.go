package repository

import (
	"context"

	"github.com/midoon/kamipa_backend/internal/domain"
	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"gorm.io/gorm"
)

type paymentRepository struct {
	simipaDB *gorm.DB
	kamipaDB *gorm.DB
}

func NewPaymentRepository(simipaDB *gorm.DB, kamipaDB *gorm.DB) domain.PaymentRepository {
	return &paymentRepository{
		simipaDB: simipaDB,
		kamipaDB: kamipaDB,
	}
}

func (r *paymentRepository) Store(ctx context.Context, p *simipa_entity.Payment) error {
	if err := r.simipaDB.WithContext(ctx).Create(p).Error; err != nil {
		return err
	}
	return nil
}

func (r *paymentRepository) StoreErrorLog(ctx context.Context, log *kamipa_entity.ErrorPaymentLog) error {
	if err := r.kamipaDB.WithContext(ctx).Create(log).Error; err != nil {
		return err
	}
	return nil
}
