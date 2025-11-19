package domain

import (
	"context"

	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/midoon/kamipa_backend/internal/model"
)

type FeeRepository interface {
	GetByStudentId(ctx context.Context, studentId int64) ([]simipa_entity.Fee, error)
	GetByFeeId(ctx context.Context, feeId int64) (simipa_entity.Fee, error)
}

type FeeUsecase interface {
	GetFees(ctx context.Context, userId string) ([]model.FeeList, error)
	GetFeeDetail(ctx context.Context, feeId int64) (model.FeeDetail, error)
}
