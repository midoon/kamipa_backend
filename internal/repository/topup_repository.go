package repository

import (
	"context"
	"time"

	"github.com/midoon/kamipa_backend/internal/domain"
	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"gorm.io/gorm"
)

type topupRepository struct {
	kamipaDB *gorm.DB
}

func NewTopupRepository(kamipaDB *gorm.DB) domain.TopupRepository {
	return &topupRepository{
		kamipaDB: kamipaDB,
	}
}

func (r *topupRepository) Save(ctx context.Context, t *kamipa_entity.Topup) error {
	return r.kamipaDB.WithContext(ctx).Create(t).Error
}

func (r *topupRepository) GetByOrderID(ctx context.Context, orderId string) (*kamipa_entity.Topup, error) {

	var t kamipa_entity.Topup

	if err := r.kamipaDB.WithContext(ctx).Where("order_id = ?", orderId).First(t).Error; err != nil {
		return &kamipa_entity.Topup{}, err
	}

	return &t, nil
}

func (r *topupRepository) UpdateStatus(ctx context.Context, orderId, status string, settlementTime *time.Time) error {
	u := map[string]interface{}{"status": status, "updated_at": time.Now()}
	if settlementTime != nil {
		u["settlement_time"] = settlementTime
	}
	return r.kamipaDB.WithContext(ctx).Model(&kamipa_entity.Topup{}).Where("order_id = ?", orderId).Updates(u).Error
}

func (r *topupRepository) SaveLog(ctx context.Context, orderId, event, status string, raw string) error {

	pl := kamipa_entity.TopupLog{
		OrderID: orderId,
		Event:   event,
		Status:  status,
		Raw:     raw,
	}
	return r.kamipaDB.WithContext(ctx).Create(&pl).Error
}
