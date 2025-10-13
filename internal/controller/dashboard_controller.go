package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
)

type DashboardController struct {
	dashboardUsecase domain.DashboardUsecase
}

func NewDashboardController(dashboardUsecase domain.DashboardUsecase) *DashboardController {
	return &DashboardController{
		dashboardUsecase: dashboardUsecase,
	}
}

func (c *DashboardController) GetNewsPosts(w http.ResponseWriter, r *http.Request) {

	postData, err := c.dashboardUsecase.GetPosts("news")
	if err != nil {
		customErr := err.(*helper.CustomError)
		helper.WriteJSON(w, customErr.Code, model.MessageResponse{
			Status:  false,
			Message: customErr.Error(),
		})
		return
	}

	resp := model.ArrayResponse[model.PostData]{
		Status:  true,
		Message: "Success",
		Data:    postData,
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}

func (c *DashboardController) GetAchievementPosts(w http.ResponseWriter, r *http.Request) {

	postData, err := c.dashboardUsecase.GetPosts("achievement")
	if err != nil {
		customErr := err.(*helper.CustomError)
		helper.WriteJSON(w, customErr.Code, model.MessageResponse{
			Status:  false,
			Message: customErr.Error(),
		})
		return
	}

	resp := model.ArrayResponse[model.PostData]{
		Status:  true,
		Message: "Success",
		Data:    postData,
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}

func (c *DashboardController) GetDetailPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	postData, err := c.dashboardUsecase.GetPostDetail(postId)
	if err != nil {
		customErr := err.(*helper.CustomError)
		helper.WriteJSON(w, customErr.Code, model.MessageResponse{
			Status:  false,
			Message: customErr.Error(),
		})
		return
	}
	resp := model.DataResponse[model.PostData]{
		Status:  true,
		Message: "Success",
		Data:    postData,
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}
