package server

import (
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/handlers"
)

type IAccountModule interface {
	Init()
}

type accountModule struct {
	*moduleFactory
	h handlers.IAccountHandler
}

func (m *moduleFactory) AccountModule() IAccountModule {

	handlers := handlers.AccountHandler()

	return &accountModule{
		moduleFactory: m,
		h:             handlers,
	}
}

func (m *accountModule) Init() {

	// handler
	m.r.Get(constants.ROUTE().ACCOUNTS, m.h.GetAllAccounts)

}
