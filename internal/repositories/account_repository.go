package repositories

import (
	"errors"
	"time"

	"github.com/goproject/internal/entities"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccountsById(accountId int) (entities.Account, error)
	FindAccounts() ([]entities.Account, error)    // To find all accounts
	AddAccount(account entities.Account) error    // To add an account
	UpdateAccount(account entities.Account) error // To update an existing account
	UpdateAmountAccountById(accountId int, ammount float64) error
	DeleteAccount(accountID entities.AccountId) error
	DeleteAllAccounts() error                                                                   // To delete an account by ID
	UpdateNameAndAmountAccount(accountID entities.AccountId, req entities.AccountRequest) error // to update only the account name
}

type accountRepository struct {
	db *gorm.DB
}

// Constructor function to create a new instance of the account repository
func AccountRepository(database *gorm.DB) IAccountRepository {
	return &accountRepository{
		db: database,
	}
}

// GetAccountsById retrieves an account by ID from the database
func (r *accountRepository) GetAccountsById(accountId int) (entities.Account, error) {
	var account entities.Account
	err := r.db.First(&account, accountId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Account{}, errors.New("account with ID %d not found")
		}
		return entities.Account{}, errors.New("error retrieving account by")
	}
	return account, nil
}

// FindAccounts retrieves all accounts from the database
func (r *accountRepository) FindAccounts() ([]entities.Account, error) {
	var accounts []entities.Account
	err := r.db.Find(&accounts).Error
	if err != nil {
		return nil, errors.New("could not find accounts: " + err.Error())
	}

	return accounts, nil
}

// AddAccount adds a new account to the database
func (r *accountRepository) AddAccount(account entities.Account) error {
	err := r.db.Create(&account).Error
	if err != nil {
		return errors.New("could not create account: " + err.Error())
	}

	return nil
}

// UpdateAccount updates an existing account in the database
func (r *accountRepository) UpdateAccount(account entities.Account) error {
	err := r.db.Model(&account).Where("account_id = ?", account.AccountID).Updates(account).Error
	if err != nil {
		return errors.New("could not update account: " + err.Error())
	}

	return nil
}

// DeleteAccount deletes an account from the database by its ID
func (r *accountRepository) DeleteAccount(accountID entities.AccountId) error {
	result := r.db.Delete(&entities.Account{}, accountID)
	if result.Error != nil {
		return errors.New("could not delete account: " + result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return errors.New("no accounts were deleted")
	}

	return nil
}

// DeleteAllAccounts deletes all accounts from the database using raw SQL
func (r *accountRepository) DeleteAllAccounts() error {
	result := r.db.Exec("DELETE FROM accounts")
	if result.Error != nil {
		return errors.New("could not delete all accounts: " + result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return errors.New("no accounts were deleted")
	}

	return nil
}

// UpdateAccountName updates only the account's name in the database
func (r *accountRepository) UpdateNameAndAmountAccount(accountID entities.AccountId, req entities.AccountRequest) error {
	err := r.db.Model(&entities.Account{}).Where("id = ?", accountID).Update("name", req.AccountName).Update("balance", req.Balance).Update("color_index", req.ColorIndex).Error
	if err != nil {
		return errors.New("could not update account name: " + err.Error())
	}

	return nil
}

func (r *accountRepository) UpdateAmountAccountById(accountId int, ammount float64) error {

	err := r.db.Model(&entities.Plan{}).Where("plan_id = ?", accountId).Update("amount", ammount).Update("update_date", time.Now()).Error
	if err != nil {
		return errors.New("could not update plan: " + err.Error())
	}

	return nil
}
