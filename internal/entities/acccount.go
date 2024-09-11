package entities

import (
	"time"
)

type AccountId = int

type AccountsList = []Account

type Color struct {
	Value int
}

type Account struct {
	AccountID      int       `gorm:"primaryKey;column:account_id" json:"accountId"` // Corresponds to account_id
	AccountName    string    `gorm:"column:name" json:"accountName"`                // Corresponds to name
	Balance        float64   `gorm:"column:amount" json:"balance"`                  // Corresponds to amount
	CreateDate     time.Time `gorm:"column:create_date" json:"createDate"`          // Corresponds to create_date
	UpdatePlanDate time.Time `gorm:"column:update_plan_date" json:"updatePlanDate"` // Corresponds to update_plan_date
	ColorIndex     int       `gorm:"column:color_index" json:"colorIndex"`          // New field
}

type AccountRequest struct {
	AccountName string  `gorm:"column:name" json:"accountName"`
	Balance     float64 `gorm:"column:amount" json:"balance"`
	ColorIndex  int     `gorm:"column:color_index" json:"colorIndex"`
}
