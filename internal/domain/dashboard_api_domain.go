package domain

import (
	"context"

	"github.com/midoon/kamipa_backend/internal/model"
)

type DashboardApiRepository interface {
	FetchPostsWithType(postType string) ([]model.PostData, error)
}

type DashboardUsecase interface {
	GetPosts(ctx context.Context, postType string) ([]model.PostData, error)
}
