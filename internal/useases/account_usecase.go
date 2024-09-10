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
		u.l.Error("cant to find accounts")
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
		u.l.Error("cant to create accounts")
	}
}
