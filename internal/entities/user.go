package entities

import "time"

type User struct {
	UserID int       `gorm:"primaryKey;autoIncrement"`
	Salary float64   `gorm:"column:salary" json:"salary"`
	Month  time.Time `gorm:"column:reset_date_planning" json:"resetDatePlanning"`
	Usages float64   `gorm:"column:current_usage_monthly" json:"currentUsageMonthly"`
}

type UserFinancialsRes struct {
	Salary float64   `gorm:"column:salary" json:"salary"`
	Month  time.Time `gorm:"column:month" json:"month"`
	Usages float64   `gorm:"column:usages" json:"usages"`
}
