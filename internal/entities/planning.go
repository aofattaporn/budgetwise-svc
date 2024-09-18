package entities

import "time"

type Plan struct {
	PlanID         int       `gorm:"primaryKey;autoIncrement"`
	Name           string    `gorm:"type:varchar(100);not null"`
	Amount         float64   `gorm:"type:decimal(10,2);not null"`
	CreateDate     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatePlanDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UserID         int       `gorm:"not null"`
	AccountID      *int      `gorm:"index"`
}

type PlanningRequest struct {
	Name      string  `gorm:"column:name" json:"name"`
	Amount    float64 `gorm:"column:amount" json:"amount"`
	UserID    int     `gorm:"column:user_id" json:"userId"`
	AccountID int     `gorm:"column:account_id" json:"accountId"`
}
