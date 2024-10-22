package useases

import (
	"time"

	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/pkg/log"
)

type IAccountUsecase interface {
	GetAllAccounts() entities.AccountsList
	CreateAccount(req entities.AccountRequest)
	DeleteAccount(accountId entities.AccountId) error
	DeleteAllAccounts() error
	PatchAccount(accountId entities.AccountId, req entities.AccountRequest)
}

type accountUsecase struct {
	r repositories.IAccountRepository
	l log.ILogger
}

func AccountUsecase(repository repositories.IAccountRepository, logger log.ILogger) IAccountUsecase {
	return &accountUsecase{
		r: repository,
		l: logger,
	}
}

func (u *accountUsecase) GetAllAccounts() entities.AccountsList {
	acccount, err := u.r.FindAccounts()
	if err != nil {
		u.l.Errorf("find accounts error %v", err)
	}
	return acccount
}

func (u *accountUsecase) CreateAccount(req entities.AccountRequest) {

	err := u.r.AddAccount(entities.Account{
		AccountName:    req.AccountName,
		Balance:        req.Balance,
		CreateDate:     time.Now(),
		UpdatePlanDate: time.Now(),
		ColorIndex:     req.ColorIndex,
		UserID:         1,
	})
	if err != nil {
		u.l.Errorf("create accounts error %v", err)
	}
}

func (u *accountUsecase) UpdateAccount(req entities.Account) {
	err := u.r.UpdateAccount(entities.Account{
		AccountID:      req.AccountID,
		AccountName:    req.AccountName,
		Balance:        req.Balance,
		CreateDate:     req.CreateDate,
		UpdatePlanDate: time.Now(),
		ColorIndex:     req.ColorIndex,
		UserID:         1,
	})
	if err != nil {
		u.l.Errorf("update accounts error %v", err)
	}
}

func (u *accountUsecase) DeleteAccount(accountId entities.AccountId) error {
	err := u.r.DeleteAccount(accountId)
	if err != nil {
		u.l.Errorf("delete accounts error %v", err)
		return err
	}
	return nil
}

func (u *accountUsecase) DeleteAllAccounts() error {
	err := u.r.DeleteAllAccounts()
	if err != nil {
		u.l.Errorf("delete accounts error %v", err)
	}
	return nil
}

func (u *accountUsecase) PatchAccount(accountId entities.AccountId, req entities.AccountRequest) {
	err := u.r.UpdateNameAndAmountAccount(accountId, req)
	if err != nil {
		u.l.Errorf("update accounts error %v", err)
	}
}
