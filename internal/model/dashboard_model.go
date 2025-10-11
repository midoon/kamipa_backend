package model

import "time"

type PostData struct {
	Id          string    `json:"Id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Type        string    `json:"type"`
	Date        time.Time `json:"date"`
}
