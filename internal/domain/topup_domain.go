package domain

import (
	"context"
	"time"

	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/model"
)

type TopupRepository interface {
	Save(ctx context.Context, t *kamipa_entity.Topup) error
	GetByOrderID(ctx context.Context, orderID string) (*kamipa_entity.Topup, error)
	UpdateStatus(ctx context.Context, orderID, status string, settlementTime *time.Time) error
	SaveLog(ctx context.Context, orderID, event, status string, raw map[string]interface{}) error
}

type TopupUsecase interface {
	CreatePayment(ctx context.Context, feeId int64, userId string) (model.TopupData, error)
	MidtransCallback(ctx context.Context, payload map[string]interface{}) error
}
