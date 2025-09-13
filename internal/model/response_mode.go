package model

type MessageResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type DataResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type PageResponse[T any] struct {
	Status   string       `json:"status"`
	Message  string       `json:"message"`
	Data     []T          `json:"data"`
	Metadata PageMetadata `json:"metadata"`
}

type PageMetadata struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`
}
