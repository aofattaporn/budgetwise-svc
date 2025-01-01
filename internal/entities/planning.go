package entities

import (
	"time"
)

type Plan struct {
	Id         int       `gorm:"primaryKey;autoIncrement"`
	Name       string    `gorm:"column:name;type:varchar(100);not null"`
	Type       string    `gorm:"column:type;type:varchar(50);not null"`
	Amount     float64   `gorm:"column:amount;type:decimal(10,2);not null"`
	Usage      float64   `gorm:"column:usages;type:decimal(10,2);not null"`
	IconIndex  int       `gorm:"column:icon_index"`
	CreateDate time.Time `gorm:"column:create_date;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdateDate time.Time `gorm:"column:update_date;type:timestamp;default:CURRENT_TIMESTAMP;on update:CURRENT_TIMESTAMP"`
	UserID     int       `gorm:"column:user_id;not null"`
	Month      time.Time `gorm:"column:month;type:date;not null"`
	AccountID  *int      `gorm:"column:account_id;default:null"`
}

type PlanDetails struct {
	Id          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Type        string    `gorm:"column:type;not null" json:"type"`
	Amount      float64   `gorm:"column:amount;not null" json:"amount"`
	Usage       float64   `gorm:"column:usages;not null" json:"usage"`
	IconIndex   int       `gorm:"column:icon_index" json:"iconIndex"`
	CreateDate  time.Time `gorm:"column:create_date;default:CURRENT_TIMESTAMP" json:"createDate"`
	UpdateDate  time.Time `gorm:"column:update_date;default:CURRENT_TIMESTAMP;on update:CURRENT_TIMESTAMP" json:"updateDate"`
	Month       time.Time `gorm:"column:month;not null" json:"month"`
	AccountName string    `gorm:"column:accountName" json:"accountName"`
}

type PlanList = []PlanDetails

type PlanningRequest struct {
	Name      string    `gorm:"column:name;not null" json:"name"`
	Type      string    `gorm:"column:type;not null" json:"type"`
	Amount    float64   `gorm:"column:amount;not null" json:"amount"`
	IconIndex int       `gorm:"column:icon_index" json:"iconIndex"`
	Month     time.Time `gorm:"column:month;not null" json:"month"`
	AccountID *int      `gorm:"column:account_id;default:null" json:"accountId"`
}
