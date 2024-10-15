// internal/entities/response.go
package entities

import "time"

// Response godoc
// @Description  Generic response format
type Response struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

// ErrorResponse godoc
// @Description  Response containing account list
type ErrorResponse struct {
	Code         int       `json:"code"`
	Timestamp    time.Time `json:"timestamp"`
	ErrorMessage string    `json:"errorMessage"`
}
