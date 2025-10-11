package usecase

import (
	"context"
	"net/http"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
)

type dashboardUsecase struct {
	dashboardApiRepository domain.DashboardApiRepository
}

func NewDashboardUsecase(dashboardApiRepository domain.DashboardApiRepository) domain.DashboardUsecase {
	return &dashboardUsecase{
		dashboardApiRepository: dashboardApiRepository,
	}
}

func (u *dashboardUsecase) GetPosts(ctx context.Context, postType string) ([]model.PostData, error) {
	posts, err := u.dashboardApiRepository.FetchPostsWithType(postType)
	if err != nil {
		return []model.PostData{}, helper.NewCustomError(http.StatusBadRequest, "Error fetch data posts", err)
	}

	return posts, nil
}
