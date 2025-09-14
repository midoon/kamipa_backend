package controller

import (
	"encoding/json"
	"net/http"

	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
)

type UserController struct {
	userUseCase domain.UserUseCase
}

func NewUserController(userUseCse domain.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCse,
	}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request := model.RegistrationUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, model.MessageResponse{
			Status:  false,
			Message: "invalid request body",
		})
		return
	}

	if err := c.userUseCase.Register(ctx, request); err != nil {
		customErr := err.(*helper.CustomError)

		helper.WriteJSON(w, customErr.Code, model.MessageResponse{
			Status:  false,
			Message: customErr.Error(),
		})
		return
	}

	helper.WriteJSON(w, http.StatusCreated, model.MessageResponse{
		Status:  true,
		Message: "registration successfully",
	})
}
