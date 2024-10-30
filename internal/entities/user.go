package entities

import "time"

type UserFinancials struct {
	UserId int       `gorm:"column:user_id;primaryKey" json:"user_id"`
	Month  time.Time `gorm:"column:month;primaryKey" json:"month"`
	Salary float64   `gorm:"column:salary" json:"salary"`
	Usages float64   `gorm:"column:usages" json:"usages"`
}

type UserFinancialsRes struct {
	Salary float64   `gorm:"column:salary" json:"salary"`
	Month  time.Time `gorm:"column:month" json:"month"`
	Usages float64   `gorm:"column:usages" json:"usages"`
}

type UserFinancialsReq struct {
	Salary float64   `gorm:"column:salary" json:"salary"`
	Month  time.Time `gorm:"column:month" json:"month"`
}
