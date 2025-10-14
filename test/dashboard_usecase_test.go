package test

import (
	"errors"
	"testing"
	"time"

	"github.com/midoon/kamipa_backend/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestGetPosts(t *testing.T) {
	t.Run("Success get all news post", func(t *testing.T) {
		deps := SetupDeps()

		expectedPosts := []model.PostData{
			{
				Id:          "1",
				Title:       "Sekolah Berjaya Menang Olimpiade Sains Nasional",
				Description: "Tim sains sekolah berhasil meraih juara 1 tingkat nasional",
				Image:       "image1.jpg",
				Type:        "news",
				Date:        time.Now(),
			},
			{
				Id:          "2",
				Title:       "Guru Berprestasi Dapat Penghargaan",
				Description: "Guru terbaik mendapat penghargaan dari dinas pendidikan",
				Image:       "image2.jpg",
				Type:        "news",
				Date:        time.Now(),
			},
		}

		// Setup mock behavior
		deps.dashboardRepo.Mock.On(
			"FetchPostsWithType",
			"news",
		).Return(expectedPosts, nil)

		// Call usecase
		result, err := deps.dashboardUsecase.GetPosts("news")

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedPosts, result)
		deps.dashboardRepo.Mock.AssertExpectations(t)
	})

	t.Run("Success get all achivement post", func(t *testing.T) {
		deps := SetupDeps()

		expectedPosts := []model.PostData{
			{
				Id:          "1",
				Title:       "Sekolah Berjaya Menang Olimpiade Sains Nasional",
				Description: "Tim sains sekolah berhasil meraih juara 1 tingkat nasional",
				Image:       "image1.jpg",
				Type:        "achivement",
				Date:        time.Now(),
			},
			{
				Id:          "2",
				Title:       "Guru Berprestasi Dapat Penghargaan",
				Description: "Guru terbaik mendapat penghargaan dari dinas pendidikan",
				Image:       "image2.jpg",
				Type:        "achivement",
				Date:        time.Now(),
			},
		}

		// Setup mock behavior
		deps.dashboardRepo.Mock.On(
			"FetchPostsWithType",
			"achivement",
		).Return(expectedPosts, nil)

		// Call usecase
		result, err := deps.dashboardUsecase.GetPosts("achivement")

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedPosts, result)
		deps.dashboardRepo.Mock.AssertExpectations(t)
	})

	t.Run("Failed to fetch posts", func(t *testing.T) {
		deps := SetupDeps()

		deps.dashboardRepo.Mock.On(
			"FetchPostsWithType",
			"news",
		).Return([]model.PostData{}, errors.New("network error"))

		result, err := deps.dashboardUsecase.GetPosts("news")

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Contains(t, err.Error(), "Error fetch data posts")

		deps.dashboardRepo.Mock.AssertExpectations(t)
	})

}

func TestGetDetailPost(t *testing.T) {
	t.Run("Success get post detail", func(t *testing.T) {
		deps := SetupDeps()

		expectedPost := model.PostData{
			Id:          "data-id-1",
			Title:       "Sekolah Berjaya Menang Olimpiade Sains Nasional",
			Description: "Tim sains sekolah berhasil meraih juara 1 tingkat nasional",
			Image:       "image1.jpg",
			Type:        "news",
			Date:        time.Now(),
		}

		// Mock repository behavior: sukses fetch detail post
		deps.dashboardRepo.Mock.On(
			"FetchDetailPost",
			"data-id-1",
		).Return(expectedPost, nil)

		// Panggil usecase
		result, err := deps.dashboardUsecase.GetPostDetail("data-id-1")

		// Assertion
		assert.NoError(t, err)
		assert.Equal(t, expectedPost, result)
		deps.dashboardRepo.Mock.AssertExpectations(t)
	})

	t.Run("Failed to fetch post detail", func(t *testing.T) {
		deps := SetupDeps()

		// Mock repository behavior: gagal fetch post detail
		deps.dashboardRepo.Mock.On(
			"FetchDetailPost",
			"data-id-404",
		).Return(model.PostData{}, errors.New("post not found"))

		// Panggil usecase
		result, err := deps.dashboardUsecase.GetPostDetail("data-id-404")

		// Assertion

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.Contains(t, err.Error(), "Error fetch detail data post")

		deps.dashboardRepo.Mock.AssertExpectations(t)
	})
}
