package repositories

import (
	"errors"

	"github.com/goproject/internal/entities"
	"gorm.io/gorm"
)

type IPlanRepository interface {
	FindAllPlans() ([]entities.Plan, error)
	AddPlan(account entities.Plan) error
}

type planRepository struct {
	db *gorm.DB
}

// Constructor function to create a new instance of the account repository
func PlanRepository(database *gorm.DB) IPlanRepository {
	return &planRepository{
		db: database,
	}
}

// FindAccounts retrieves all plans from the database
func (r *planRepository) FindAllPlans() ([]entities.Plan, error) {
	return []entities.Plan{}, nil
}

// AddAccount adds a new plan to the database
func (r *planRepository) AddPlan(account entities.Plan) error {
	err := r.db.Create(&account).Error
	if err != nil {
		return errors.New("could not create account: " + err.Error())
	}

	return nil
}
