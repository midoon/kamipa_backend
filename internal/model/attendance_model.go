package model

import "time"

type AttendanceData struct {
	StudentId int64     `json:"student_id"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
}
