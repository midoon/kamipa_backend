package controller

import (
	"encoding/json"
	"net/http"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
)

type TopupController struct {
	topupUsease domain.TopupUsecase
}

func NewTopupController(topupUsecase domain.TopupUsecase) *TopupController {
	return &TopupController{
		topupUsease: topupUsecase,
	}
}

func (c *TopupController) CreatePayment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.Context().Value(helper.UserIDKey).(string)
	request := model.CreateTopupRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, model.MessageResponse{
			Status:  false,
			Message: "invalid request body",
		})
		return
	}

	topupData, err := c.topupUsease.CreatePayment(ctx, request.FeeId, userId)

	if err != nil {
		customErr := err.(*helper.CustomError)

		helper.WriteJSON(w, customErr.Code, model.MessageResponse{
			Status:  false,
			Message: customErr.Error(),
		})
		return
	}

	resp := model.DataResponse[model.TopupData]{
		Status:  true,
		Message: "Success get snap url",
		Data:    topupData,
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}

func (c *TopupController) CallbakcHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse JSON body
	var callback map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&callback); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, model.MessageResponse{
			Status:  false,
			Message: "invalid callback payload",
		})
		return
	}

	// Process callback in usecase
	err := c.topupUsease.MidtransCallback(ctx, callback)
	if err != nil {
		// Jika error berasal dari CustomError
		if customErr, ok := err.(*helper.CustomError); ok {
			helper.WriteJSON(w, customErr.Code, model.MessageResponse{
				Status:  false,
				Message: customErr.Error(),
			})
			return
		}

		// Fallback jika bukan custom error
		helper.WriteJSON(w, http.StatusInternalServerError, model.MessageResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	// Success response
	helper.WriteJSON(w, http.StatusOK, model.MessageResponse{
		Status:  true,
		Message: "callback processed successfully",
	})
}
