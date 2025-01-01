package repositories

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
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
	// Ensure the month field is set to the first day of the month
	// req.Month = req.Month.UTC().Truncate(24 * time.Hour)

	err := r.db.Create(&req).Error
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return fmt.Errorf("duplicate entry: %w", err)
		}
		return err
	}
	return nil
}
