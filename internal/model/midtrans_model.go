package model

type CreateTopupRequest struct {
	FeeId int64 `json:"fee_id"`
}

type TopupData struct {
	OrderID     string `json:"order_id"`
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}
