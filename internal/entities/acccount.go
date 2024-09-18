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
	AccountID      int       `gorm:"primaryKey;column:account_id" json:"accountId"`
	AccountName    string    `gorm:"column:name" json:"accountName"`
	Balance        float64   `gorm:"column:amount" json:"balance"`
	CreateDate     time.Time `gorm:"column:create_date" json:"createDate"`
	UpdatePlanDate time.Time `gorm:"column:update_plan_date" json:"updatePlanDate"`
	ColorIndex     int       `gorm:"column:color_index" json:"colorIndex"`
	UserID         int       `gorm:"column:user_id" json:"userId"`
}

type AccountRequest struct {
	AccountName string  `gorm:"column:name" json:"accountName"`
	Balance     float64 `gorm:"column:amount" json:"balance"`
	ColorIndex  int     `gorm:"column:color_index" json:"colorIndex"`
}
