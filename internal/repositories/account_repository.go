package repositories

import (
	"errors"

	"github.com/goproject/internal/entities"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	FindAccounts() ([]entities.Account, error)
	AddAccount(account entities.Account) error
}

type accountRepository struct {
	db *gorm.DB
}

func AccountRepository(database *gorm.DB) IAccountRepository {
	return &accountRepository{
		db: database,
	}
}

func (r *accountRepository) FindAccounts() ([]entities.Account, error) {
	var accounts []entities.Account
	err := r.db.Find(&accounts).Error
	if err != nil {
		return nil, errors.New("could not find accounts: " + err.Error())
	}

	return accounts, nil
}

func (r *accountRepository) AddAccount(account entities.Account) error {
	err := r.db.Create(&account).Error
	if err != nil {
		return errors.New("could not create account: " + err.Error())
	}

	return nil
}
