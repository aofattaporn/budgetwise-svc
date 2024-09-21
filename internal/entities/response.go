package entities

import "time"

type Response struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

type ErrorResponse struct {
	Code         int       `json:"code"`
	Timestamp    time.Time `json:"timestamp"`
	ErrorMessage string    `json:"errorMessage"`
}
