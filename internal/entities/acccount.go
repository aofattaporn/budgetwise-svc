package entities

import (
	"time"
)

type AccountId = int

type AccountsList = []Account

type Account struct {
	AccountID      int       `json:"account_id" db:"account_id"`
	Name           string    `json:"name" db:"name"`
	Amount         float64   `json:"amount" db:"amount"`
	CreateDate     time.Time `json:"create_date" db:"create_date"`
	UpdatePlanDate time.Time `json:"update_plan_date" db:"update_plan_date"`
}

type AccountRequest struct {
	Name   string  `json:"name" db:"name"`
	Amount float64 `json:"amount" db:"amount"`
}
