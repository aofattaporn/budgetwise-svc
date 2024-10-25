package entities

import (
	"time"
)

type Plan struct {
	PlanID         int       `gorm:"primaryKey;autoIncrement"`
	Name           string    `gorm:"type:varchar(100);not null"`
	Amount         float64   `gorm:"type:decimal(10,2);not null"`
	IconIndex      int       `gorm:"column:icon_index" json:"iconIndex"`
	CreateDate     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatePlanDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UserID         int       `gorm:"not null"`
	AccountID      int       `gorm:"index"`
	Usage          float64   `gorm:"column:usages" json:"usages"`
}

type PlanDetails struct {
	PlanID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string    `gorm:"column:name" json:"name"`
	Usage          float64   `gorm:"column:usages" json:"usages"`
	Amount         float64   `gorm:"column:amount" json:"amount"`
	CreateDate     time.Time `gorm:"column:create_date" json:"createDate"`
	UpdatePlanDate time.Time `gorm:"column:update_date" json:"updateDate"`
	AccountName    string    `gorm:"column:accountName" json:"accountName"`
	IconIndex      int       `gorm:"column:icon_index" json:"iconIndex"`
}

type PlanList = []PlanDetails

type PlanningRequest struct {
	Name      string  `gorm:"column:name" json:"name"`
	Amount    float64 `gorm:"column:amount" json:"amount"`
	AccountID int     `gorm:"column:account_id" json:"accountId"`
	IconIndex int     `gorm:"column:icon_index" json:"iconIndex"`
}
