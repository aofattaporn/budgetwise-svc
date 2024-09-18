package entities

import "time"

type User struct {
	UserID            int       `gorm:"primaryKey;autoIncrement"`
	Salary            float64   `gorm:"column:salary" json:"salary"`
	ResetDatePlanning time.Time `gorm:"column:reset_date_planning" json:"resetDatePlanning"`
}

type SalaryAndResetDate struct {
	Salary            float64   `gorm:"column:salary" json:"salary"`
	ResetDatePlanning time.Time `gorm:"column:reset_date_planning" json:"resetDatePlanning"`
}
