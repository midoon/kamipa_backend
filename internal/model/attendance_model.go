package model

import "time"

type AttendanceData struct {
	StudentId int64     `json:"student_id"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
	Activity  string    `json:"activity"`
}

type AttendanceSummary struct {
	Activity    string    `json:"activity"`
	StartDate   time.Time `json:"start_date"`
	CurrentDate time.Time `json:"current_date"`
	Hadir       int16     `json:"hadir"`
	Izin        int16     `json:"izin"`
	Sakit       int16     `json:"sakit"`
	Alpha       int16     `json:"alpha"`
}
