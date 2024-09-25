package repositories

import (
	"errors"
	"time"

	"github.com/goproject/internal/entities"
	"gorm.io/gorm"
)

type IPlanRepository interface {
	FindAllPlans() ([]entities.PlanDetails, error)
	AddPlan(account entities.Plan) error
	UpdatePlan(account entities.Plan) error
	DeletePlanById(planId int) error
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

type Account struct {
	AccountID      int       `gorm:"primaryKey;column:account_id" json:"accountId"`
	AccountName    string    `gorm:"column:name" json:"accountName"`
	Balance        float64   `gorm:"column:amount" json:"balance"`
	CreateDate     time.Time `gorm:"column:create_date" json:"createDate"`
	UpdatePlanDate time.Time `gorm:"column:update_plan_date" json:"updatePlanDate"`
	ColorIndex     int       `gorm:"column:color_index" json:"colorIndex"`
	UserID         int       `gorm:"column:user_id" json:"userId"`
}

// FindAllPlans retrieves all plans with their associated accounts from the database
func (r *planRepository) FindAllPlans() ([]entities.PlanDetails, error) {
	var plans []entities.PlanDetails
	err := r.db.Model(&entities.Plan{}).
		Select(
			"plans.plan_id AS plan_id, " +
				"plans.name AS name, " +
				"plans.plan_usage AS plan_usage, " +
				"plans.amount AS amount, " +
				"plans.create_date AS create_date, " +
				"plans.icon_index AS icon_index, " +
				"plans.update_plan_date AS update_plan_date, " +
				"accounts.name AS accountName").
		Joins("LEFT JOIN accounts ON plans.account_id = accounts.account_id").
		Scan(&plans).Error

	if err != nil {
		return nil, errors.New("could not find plans: " + err.Error())
	}
	return plans, nil
}

// addAccount adds a new plan to the database
func (r *planRepository) AddPlan(plan entities.Plan) error {
	if err := r.db.First(&entities.Account{}, plan.AccountID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("referenced record not found")
		}
		return errors.New("could not check account: " + err.Error())
	}

	err := r.db.Create(&plan).Error
	if err != nil {
		return errors.New("could not create plan: " + err.Error())
	}

	return nil
}

// addAccount adds a new plan to the database
func (r *planRepository) UpdatePlan(plan entities.Plan) error {
	if err := r.db.First(&entities.Account{}, plan.AccountID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("referenced record not found")
		}
		return errors.New("could not check account: " + err.Error())
	}

	err := r.db.Model(&entities.Plan{}).Where("plan_id = ?", plan.PlanID).Update("name", plan.Name).Update("amount", plan.Amount).Update("icon_index", plan.IconIndex).Update("account_id", plan.AccountID).Update("update_plan_date", plan.UpdatePlanDate).Error
	if err != nil {
		return errors.New("could not update plan: " + err.Error())
	}

	return nil
}

// DeleteAccount deletes an account from the database by its ID
func (r *planRepository) DeletePlanById(planId int) error {
	err := r.db.Delete(&entities.Plan{}, planId).Error
	if err != nil {
		return errors.New("could not delete plan: " + err.Error())
	}
	return nil
}
