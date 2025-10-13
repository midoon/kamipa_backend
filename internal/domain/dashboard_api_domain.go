package domain

import (
	"github.com/midoon/kamipa_backend/internal/model"
)

type DashboardApiRepository interface {
	FetchPostsWithType(postType string) ([]model.PostData, error)
	FetchDetailPost(postId string) (model.PostData, error)
}

type DashboardUsecase interface {
	GetPosts(postType string) ([]model.PostData, error)
	GetPostDetail(postId string) (model.PostData, error)
}
