package mockrepo

import (
	"context"

	"github.com/midoon/kamipa_backend/internal/entity/simipa_entity"
	"github.com/stretchr/testify/mock"
)

type StudentRepositoryMock struct {
	Mock mock.Mock
}

func (r *StudentRepositoryMock) GetByNisn(ctx context.Context, nisn string) (simipa_entity.Student, error) {
	arguments := r.Mock.Called(ctx, nisn)

	// ambil student
	student, _ := arguments.Get(0).(simipa_entity.Student)

	// ambil error
	err, _ := arguments.Get(1).(error)

	return student, err
}

func (r *StudentRepositoryMock) CountByNisn(ctx context.Context, nisn string) (int16, error) {
	args := r.Mock.Called(ctx, nisn)

	count, _ := args.Get(0).(int16)
	err, _ := args.Get(1).(error)

	return count, err
}
