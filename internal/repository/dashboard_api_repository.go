package repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/model"
)

type dashboardApiRepository struct {
	client  *http.Client
	baseUrl string
}

func NewDashboardApiRepository(client *http.Client, baseUrl string) domain.DashboardApiRepository {
	return &dashboardApiRepository{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (r *dashboardApiRepository) FetchPostsWithType(postType string) ([]model.PostData, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/posts", r.baseUrl), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("type", postType)
	q.Add("limit", "5")
	req.URL.RawQuery = q.Encode()

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	apiResp := model.ArrayResponse[model.PostData]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Data, nil
}

func (r *dashboardApiRepository) FetchDetailPost(postId string) (model.PostData, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/posts/%s", r.baseUrl, postId), nil)
	if err != nil {
		return model.PostData{}, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return model.PostData{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.PostData{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	apiResp := model.DataResponse[model.PostData]{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return model.PostData{}, err
	}

	return apiResp.Data, nil

}
