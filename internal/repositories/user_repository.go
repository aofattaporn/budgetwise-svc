package repositories

import (
	"fmt"

	"github.com/goproject/internal/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetSalaryAndDateReset(userID int, monthYear string) (*entities.UserFinancialsRes, error)
	AddNewSalaryBymonth(req *entities.UserFinancials) error
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

func (r *userRepository) GetSalaryAndDateReset(userID int, monthYear string) (*entities.UserFinancialsRes, error) {
	var user entities.UserFinancialsRes
	err := r.db.Table("user_financials").Select("salary", "month", "usages").
		Where("user_id = ? AND DATE(month) = ?", userID, monthYear+"-01").
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) AddNewSalaryBymonth(req *entities.UserFinancials) error {
	err := r.db.Create(&req).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return fmt.Errorf("duplicate entry: %w", err)
		}
		return err
	}
	return nil
}
