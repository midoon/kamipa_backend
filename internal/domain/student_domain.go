package domain

import (
	"context"

	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
)

type StudentRepository interface {
	GetByNisn(ctx context.Context, nisn string) (simipa_entity.Student, error)
	CountByNisn(ctx context.Context, nisn string) (int16, error)
}
