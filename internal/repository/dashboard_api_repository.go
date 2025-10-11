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
