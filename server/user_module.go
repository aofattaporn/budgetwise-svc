package server

import (
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/handlers"
)

type IUserModule interface {
	Init()
}

type userModuleModule struct {
	*moduleFactory
	h handlers.IUserHandler
}

func (m *moduleFactory) UserModule() IUserModule {

	handlers := handlers.UserHandler()

	return &userModuleModule{
		moduleFactory: m,
		h:             handlers,
	}
}

func (m *userModuleModule) Init() {

	// handlers
	m.r.Get(constants.ROUTE().SALARY, m.h.GetSalary)

}
