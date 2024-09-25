package entities

import "time"

type Transaction struct {
	TransactionId int       `gorm:"primaryKey;column:transaction_id" json:"transactionId"`
	Name          string    `gorm:"column:name" json:"name"`
	Amount        float64   `gorm:"column:amount" json:"amount"`
	Operation     string    `gorm:"column:operation" json:"operation"`
	CreateDate    time.Time `gorm:"column:create_date" json:"createDate"`
	UpdateDate    time.Time `gorm:"column:update_date" json:"updateDate"`
	UserID        int       `gorm:"column:user_id" json:"userId"`
	PlanId        int       `gorm:"column:plan_id" json:"planId"`
	AccountId     int       `gorm:"column:account_id" json:"accountId"`
}

type TransactionReq struct {
	Name      string  `gorm:"column:name" json:"name"`
	Amount    float64 `gorm:"column:amount" json:"amount"`
	Operation string  `gorm:"column:operation" json:"operation"`
	PlanId    int     `gorm:"column:plan_id" json:"planId"`
	AccountId int     `gorm:"column:account_id" json:"accountId"`
}

type TransactionListRes = []TransactionRes

type TransactionRes struct {
	TransactionId int       `gorm:"primaryKey;column:transaction_id" json:"transactionId"`
	Name          string    `gorm:"column:name" json:"name"`
	Amount        float64   `gorm:"column:amount" json:"amount"`
	Operation     string    `gorm:"column:operation" json:"operation"`
	CreateDate    time.Time `gorm:"column:create_date" json:"createDate"`
	UpdateDate    int       `gorm:"column:update_date" json:"updateDate"`
	UserID        int       `gorm:"column:user_id" json:"userId"`
	PlanName      string    `gorm:"column:plan_name" json:"planName"`
	AccountName   string    `gorm:"column:account_name" json:"accountName"`
}
