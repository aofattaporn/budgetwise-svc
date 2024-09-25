package repositories

import (
	"errors"
	"time"

	"github.com/goproject/internal/entities"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	FindListTransaction() (entities.TransactionListRes, error)
	AddTransaction(account entities.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

// Constructor function to create a new instance of the account repository
func TransactionRepository(database *gorm.DB) ITransactionRepository {
	return &transactionRepository{
		db: database,
	}
}

// FindAllPlans retrieves all plans with their associated accounts from the database
func (r *transactionRepository) FindListTransaction() ([]entities.TransactionRes, error) {
	var transactions []entities.TransactionRes
	today := time.Now().Format("YYYY-MM-DD")

	// Filter transactions based on today's date
	err := r.db.Where("DATE(create_date) = ?", today).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// Creat Transaction
func (r *transactionRepository) AddTransaction(account entities.Transaction) error {
	err := r.db.Create(&account).Error
	if err != nil {
		return errors.New("could not create transaction: " + err.Error())
	}
	return nil
}
