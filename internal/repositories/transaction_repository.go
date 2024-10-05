package repositories

import (
	"errors"
	"strings"

	"github.com/goproject/internal/customerrors"
	"github.com/goproject/internal/entities"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	FindListTransaction(date string) (entities.TransactionListRes, error)
	AddTransaction(account entities.Transaction) error
	DeleteTransactionById(id int) error
	TruncateTransaction() error
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

func (r *transactionRepository) FindListTransaction(date string) ([]entities.TransactionRes, error) {

	var transactions []entities.TransactionRes

	// Trim any potential whitespace from the date parameter
	trimmedDate := strings.TrimSpace(date)

	// Query transactions based on the passed date, joining plans and accounts
	err := r.db.Model(entities.Transaction{}).
		Where("DATE(transactions.create_date) = ?", trimmedDate). // Dynamically use the date parameter
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
		Joins("LEFT JOIN accounts ON transactions.account_id = accounts.account_id"). // Corrected join to use `transactions.account_id`
		Scan(&transactions).Error

	if err != nil {
		return nil, errors.New("[database errors]: " + err.Error())
	}

	return transactions, nil
}

func (r *transactionRepository) AddTransaction(transaction entities.Transaction) error {
	err := r.db.Create(&transaction).Error

	if err != nil {
		if strings.Contains(err.Error(), "foreign key constraint fails") {
			return customerrors.FOREIGN_KEY_VIOLATION_ERROR("Account ID or Plan ID does not exist")
		}
		return errors.New("[database errors]: " + err.Error())
	}

	return nil
}

func (r *transactionRepository) DeleteTransactionById(id int) error {
	err := r.db.Delete(&entities.Transaction{}, id).Error
	if err != nil {
		return errors.New("[database errors]: " + err.Error())
	}
	return nil
}

func (r *transactionRepository) TruncateTransaction() error {
	result := r.db.Exec("DELETE FROM transactions")
	if result.Error != nil {
		return errors.New("[database errors]: " + result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return errors.New("[database errors]: no accounts were deleted")
	}

	return nil
}
