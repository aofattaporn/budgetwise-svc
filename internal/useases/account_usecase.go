package useases

import (
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/repositories"
)

type IAccountUsecase interface {
	GetAllAccounts() entities.Account
	CreateAccount(req entities.AccountRequest)
}

type accountUsecase struct {
	r repositories.IAccountRepository
}

func AccountUsecase(repository repositories.IAccountRepository) IAccountUsecase {
	return &accountUsecase{
		r: repository,
	}
}

func (u *accountUsecase) GetAllAccounts() entities.Account {
	return entities.Account{}
}

func (u *accountUsecase) CreateAccount(req entities.AccountRequest) {
}
