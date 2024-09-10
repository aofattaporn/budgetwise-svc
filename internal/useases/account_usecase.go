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
	UpdateAccount(req entities.Account)
	DeleteAccount(accountId entities.AccountId) error
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
		Name:           req.Name,
		Amount:         req.Amount,
		CreateDate:     time.Now(),
		UpdatePlanDate: time.Now(),
	})
	if err != nil {
		u.l.Errorf("create accounts error %v", err)
	}
}

func (u *accountUsecase) UpdateAccount(req entities.Account) {
	err := u.r.UpdateAccount(entities.Account{
		AccountID:      req.AccountID,
		Name:           req.Name,
		Amount:         req.Amount,
		CreateDate:     req.CreateDate,
		UpdatePlanDate: time.Now(),
	})
	if err != nil {
		u.l.Errorf("update accounts error %v", err)
	}
}

func (u *accountUsecase) DeleteAccount(accountId entities.AccountId) error {
	err := u.r.DeleteAccount(accountId)
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
