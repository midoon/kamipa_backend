package usecase

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/midoon/kamipa_backend/internal/domain"
	kamipa_entity "github.com/midoon/kamipa_backend/internal/entity/kamipa_entitiy"
	"github.com/midoon/kamipa_backend/internal/helper"
	"github.com/midoon/kamipa_backend/internal/model"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type topupUsecase struct {
	midtransKey       string
	midtransEnv       bool
	topupRespository  domain.TopupRepository
	feeRepository     domain.FeeRepository
	userRepository    domain.UserRepository
	studentRepository domain.StudentRepository
}

func NewTopupUsecase(midtransKey string, midtransEnv bool, topupRepository domain.TopupRepository, feeRepository domain.FeeRepository, userRepository domain.UserRepository, studentRepository domain.StudentRepository) domain.TopupUsecase {
	return &topupUsecase{
		midtransKey:       midtransKey,
		midtransEnv:       midtransEnv,
		topupRespository:  topupRepository,
		feeRepository:     feeRepository,
		userRepository:    userRepository,
		studentRepository: studentRepository,
	}
}

func (u *topupUsecase) CreatePayment(ctx context.Context, feeId int64, userId string) (model.TopupData, error) {
	snapClient := snap.Client{}
	isProdMidtrans := u.midtransEnv

	if isProdMidtrans {
		snapClient.New(u.midtransKey, midtrans.Production)
	} else {
		snapClient.New(u.midtransKey, midtrans.Sandbox)
	}

	id := uuid.New().String()
	orderId := "FEE-" + id

	fee, err := u.feeRepository.GetByFeeId(ctx, feeId)
	if err != nil {
		return model.TopupData{}, helper.NewCustomError(http.StatusInternalServerError, "Error get fee data", err)
	}

	user, err := u.userRepository.GetById(ctx, userId)
	if err != nil {
		return model.TopupData{}, helper.NewCustomError(http.StatusInternalServerError, "Error get user data", err)
	}

	amount := int64(fee.Amount) - int64(fee.PaidAmount)
	if amount == 0 {
		return model.TopupData{}, helper.NewCustomError(http.StatusBadRequest, "Fee totally paid", nil)
	}

	// cek apakah fee itu merupakan fee dari user
	student, err := u.studentRepository.GetByNisn(ctx, user.StudentNisn)
	if err != nil {
		return model.TopupData{}, helper.NewCustomError(http.StatusInternalServerError, "Error get student data", err)
	}

	if student.ID != fee.StudentId {
		return model.TopupData{}, helper.NewCustomError(http.StatusBadRequest, "Payment declined, This bill is not yours", err)
	}

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: amount,
		},
		CreditCard: &snap.CreditCardDetails{Secure: true},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.ID,
			Email: user.Email,
			Phone: user.Phone,
		},
		Items: &[]midtrans.ItemDetails{{
			ID:    "FEE-" + strconv.FormatInt(feeId, 10),
			Qty:   1,
			Price: amount,
			Name:  fee.PaymentType.Name,
		}},
	}

	resp, _ := snapClient.CreateTransaction(snapReq)

	if resp == nil {
		return model.TopupData{}, helper.NewCustomError(http.StatusInternalServerError, "Error generate snap url", nil)
	}

	now := time.Now()
	expiry := now.Add(24 * time.Hour)

	t := &kamipa_entity.Topup{
		OrderID:         orderId,
		FeeID:           feeId,
		UserID:          userId,
		Amount:          amount,
		SnapToken:       resp.Token,
		SnapTokenExpiry: &expiry,
		Status:          "pending",
		RawResponse:     "",
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	if err := u.topupRespository.Save(ctx, t); err != nil {
		return model.TopupData{}, helper.NewCustomError(http.StatusInternalServerError, "Error store data topup", err)
	}

	// save create log
	raw, err := json.Marshal(map[string]interface{}{"snap_response": resp})
	if err != nil {
		return model.TopupData{}, helper.NewCustomError(http.StatusInternalServerError, "Error marshal data", err)
	}
	if err := u.topupRespository.SaveLog(ctx, orderId, "create", "pending", string(raw)); err != nil {
		return model.TopupData{}, helper.NewCustomError(http.StatusInternalServerError, "Error store data log topup", err)
	}

	topupData := &model.TopupData{
		OrderID:     orderId,
		Token:       resp.Token,
		RedirectURL: resp.RedirectURL,
	}

	return *topupData, nil
}

func (u *topupUsecase) MidtransCallback(ctx context.Context, payload map[string]interface{}) error {
	orderId, _ := payload["order_id"].(string)
	status, _ := payload["transaction_status"].(string)

	// persist raw callback
	rawBytes, err := json.Marshal(payload)
	if err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, "Error marshal data", err)
	}

	if err := u.topupRespository.SaveLog(ctx, orderId, "callback", status, string(rawBytes)); err != nil {
		return helper.NewCustomError(http.StatusInternalServerError, "Error store data log topup", err)
	}

	// map midtrans status to internal
	switch status {
	case "settlement":
		t := time.Now()
		return u.topupRespository.UpdateStatus(ctx, orderId, "paid", &t)
	case "pending":
		return u.topupRespository.UpdateStatus(ctx, orderId, "pending", nil)
	case "expire":
		return u.topupRespository.UpdateStatus(ctx, orderId, "expired", nil)
	case "deny", "cancel":
		return u.topupRespository.UpdateStatus(ctx, orderId, "failed", nil)
	default:
		// unknown status: still save log
		return nil
	}
}
