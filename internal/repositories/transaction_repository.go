package repositories

import (
	"errors"
	"fmt"
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

func (r *transactionRepository) FindListTransaction() ([]entities.TransactionRes, error) {

	var transactions []entities.TransactionRes
	today := time.Now()

	// Filter transactions based on today's date and join plans and accounts
	err := r.db.Model(entities.Transaction{}).
		Where("DATE(transactions.create_date) = ?", today.Format("2006-01-02")). // Format date without time
		Select(
			"transactions.transaction_id AS transaction_id, " +
				"transactions.name AS name, " +
				"transactions.amount AS amount, " +
				"transactions.operation AS operation, " +
				"transactions.create_date AS create_date, " +
				"transactions.update_date AS update_date, " +
				"plans.icon_index AS icon_index, " +
				"plans.name AS plan_name, " + // Plan name
				"accounts.name AS account_name"). // Account name
		Joins("LEFT JOIN plans ON transactions.plan_id = plans.plan_id").
		Joins("LEFT JOIN accounts ON plans.account_id = accounts.account_id").
		Scan(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// Creat Transaction
func (r *transactionRepository) AddTransaction(transaction entities.Transaction) error {

	fmt.Println(transaction)
	fmt.Println("====================")

	err := r.db.Create(&transaction).Error
	if err != nil {
		return errors.New("could not create transaction: " + err.Error())
	}
	return nil
}
