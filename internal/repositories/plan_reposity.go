package repositories

import (
	"errors"
	"time"

	"github.com/goproject/internal/entities"
	"gorm.io/gorm"
)

type IPlanRepository interface {
	GetPlanById(planId int) (entities.Plan, error)
	FindAllPlans(monthYear string) ([]entities.PlanDetails, error)
	AddPlan(account entities.Plan) error
	UpdatePlan(account entities.Plan) error
	UpdateAmountPlanById(planId int, amount float64) error
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
	AccountID   int       `gorm:"primaryKey;column:account_id" json:"accountId"`
	AccountName string    `gorm:"column:name" json:"accountName"`
	Balance     float64   `gorm:"column:amount" json:"balance"`
	CreateDate  time.Time `gorm:"column:create_date" json:"createDate"`
	UpdateDate  time.Time `gorm:"column:update_date" json:"updateDate"`
	ColorIndex  int       `gorm:"column:color_index" json:"colorIndex"`
	UserID      int       `gorm:"column:user_id" json:"userId"`
}

// GetAccountsById retrieves an account by ID from the database
func (r *planRepository) GetPlanById(planId int) (entities.Plan, error) {
	var plan entities.Plan
	err := r.db.First(&plan, planId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Plan{}, errors.New("account with ID %d not found")
		}
		return entities.Plan{}, errors.New("error retrieving account by")
	}
	return plan, nil
}

// FindAllPlans retrieves all plans with their associated accounts from the database
func (r *planRepository) FindAllPlans(monthYear string) ([]entities.PlanDetails, error) {
	var plans []entities.PlanDetails
	err := r.db.Model(&entities.Plan{}).
		Select(
			"plans.id AS id, "+
				"plans.name AS name, "+
				"plans.type AS type, "+
				"plans.usages AS usages, "+
				"plans.amount AS amount, "+
				"plans.create_date AS create_date, "+
				"plans.icon_index AS icon_index, "+
				"plans.update_date AS update_date, "+
				"plans.month AS month, "+
				"accounts.name AS accountName").
		Joins("LEFT JOIN accounts ON plans.account_id = accounts.id").
		Where("DATE(plans.month) = ?", monthYear+"-01").
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

	err := r.db.Model(&entities.Plan{}).Where("id = ?", plan.Id).Update("name", plan.Name).Update("amount", plan.Amount).Update("icon_index", plan.IconIndex).Update("account_id", plan.AccountID).Error
	if err != nil {
		return errors.New("could not update plan: " + err.Error())
	}

	return nil
}

func (r *planRepository) UpdateAmountPlanById(planId int, ammount float64) error {
	err := r.db.Model(&entities.Plan{}).Where("id = ?", planId).Update("amount", ammount).Update("update_date", time.Now()).Error
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
