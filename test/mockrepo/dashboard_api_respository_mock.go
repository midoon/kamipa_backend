package mockrepo

import (
	"github.com/midoon/kamipa_backend/internal/model"
	"github.com/stretchr/testify/mock"
)

type DashboardRepositoryMock struct {
	Mock mock.Mock
}

func (r *DashboardRepositoryMock) FetchPostsWithType(postType string) ([]model.PostData, error) {
	args := r.Mock.Called(postType)

	posts, _ := args.Get(0).([]model.PostData)
	err := args.Error(1)

	return posts, err
}

func (r *DashboardRepositoryMock) FetchDetailPost(postId string) (model.PostData, error) {
	args := r.Mock.Called(postId)

	post, _ := args.Get(0).(model.PostData)
	err := args.Error(1)

	return post, err
}
