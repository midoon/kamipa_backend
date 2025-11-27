package domain

import (
	"context"

	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
)

type PaymentRepository interface {
	Store(ctx context.Context, p *simipa_entity.Payment) error
	StoreErrorLog(ctx context.Context, log *kamipa_entity.ErrorPaymentLog) error
}
