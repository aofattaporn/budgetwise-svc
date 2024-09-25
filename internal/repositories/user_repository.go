package repositories

import (
	"github.com/goproject/internal/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetSalaryAndDateReset(userID int) (*entities.SalaryAndResetDate, error)
}

type userRepository struct {
	db *gorm.DB
}

// Constructor function to create a new instance of the account repository
func UserRepository(database *gorm.DB) IUserRepository {
	return &userRepository{
		db: database,
	}
}

func (r *userRepository) GetSalaryAndDateReset(userID int) (*entities.SalaryAndResetDate, error) {
	var user entities.SalaryAndResetDate
	err := r.db.Table("users").Select("salary", "reset_date_planning", "current_usage_monthly").Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
