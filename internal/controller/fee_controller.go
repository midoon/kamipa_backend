package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/midoon/kamipa_backend/internal/domain"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
)

type FeeController struct {
	feeUsecase domain.FeeUsecase
}

func NewFeeController(feeUsecase domain.FeeUsecase) *FeeController {
	return &FeeController{
		feeUsecase: feeUsecase,
	}
}

func (c *FeeController) GetFees(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.Context().Value(helper.UserIDKey).(string)

	fees, err := c.feeUsecase.GetFees(ctx, userId)
	if err != nil {
		customErr := err.(*helper.CustomError)
		helper.WriteJSON(w, customErr.Code, model.MessageResponse{
			Status:  false,
			Message: customErr.Error(),
		})
		return
	}

	resp := model.ArrayResponse[model.FeeList]{
		Status:  true,
		Message: "Success get fee data",
		Data:    fees,
	}

	helper.WriteJSON(w, http.StatusOK, resp)
}

func (c *FeeController) GetFeeDetail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	feeIDStr := vars["feeId"]

	feeId, err := strconv.ParseInt(feeIDStr, 10, 64)
	if err != nil {
		helper.WriteJSON(w, http.StatusBadRequest, model.MessageResponse{
			Status:  false,
			Message: "Invalid fee ID",
		})
		return
	}
	feeDetail, err := c.feeUsecase.GetFeeDetail(ctx, feeId)
	if err != nil {
		customErr := err.(*helper.CustomError)
		helper.WriteJSON(w, customErr.Code, model.MessageResponse{
			Status:  false,
			Message: customErr.Message,
		})
		return
	}

	resp := model.DataResponse[model.FeeDetail]{
		Status:  true,
		Message: "Success get fee detail data",
		Data:    feeDetail,
	}

	helper.WriteJSON(w, http.StatusOK, resp)

}
