package mockrepo

import (
	"context"
	"net/http"

	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/stretchr/testify/mock"
)

type StudentRepositoryMock struct {
	Mock mock.Mock
}

func (r *StudentRepositoryMock) GetByNisn(ctx context.Context, nisn string) (simipa_entity.Student, error) {
	arguments := r.Mock.Called(ctx, nisn)
	// jika data yang dikiramkan berupa nill
	if arguments.Get(0) == nil {
		return simipa_entity.Student{}, helper.NewCustomError(http.StatusInternalServerError, "failed to store user", nil)
	}

	return simipa_entity.Student{}, nil
}
