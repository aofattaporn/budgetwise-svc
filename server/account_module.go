package server

import (
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/handlers"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/internal/useases"
)

type IAccountModule interface {
	Init()
}

type accountModule struct {
	*moduleFactory
	h handlers.IAccountHandler
}

func (m *moduleFactory) AccountModule() IAccountModule {

	repository := repositories.AccountRepository(m.s.db.GetDb())
	usecase := useases.AccountUsecase(repository, m.s.logger)
	handlers := handlers.AccountHandler(usecase)

	return &accountModule{
		moduleFactory: m,
		h:             handlers,
	}
}

func (m *accountModule) Init() {

	// handler
	m.r.Get(constants.ROUTE().ACCOUNTS, m.h.GetAllAccounts)
	m.r.Post(constants.ROUTE().ACCOUNTS, m.h.CreateAccount)

}
