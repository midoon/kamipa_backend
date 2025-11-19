package model

import "time"

type FeeList struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type FeeDetail struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Amount          float64   `json:"amount"`
	Status          string    `json:"status"`
	RemainingAmount float64   `json:"remaining_amount"`
	DueDate         time.Time `json:"due_date"`
}
