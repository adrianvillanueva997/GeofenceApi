package models

import "time"

type ErrorMonitor struct {
	Date  time.Time `json:"date"`
	Error error     `json:"error"`
}
