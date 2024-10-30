package server

import (
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/handlers"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/internal/useases"
)

type IUserModule interface {
	Init()
}

type userModuleModule struct {
	*moduleFactory
	h handlers.IUserHandler
}

func (m *moduleFactory) UserModule() IUserModule {

	repository := repositories.UserRepository(m.s.db.GetDb())
	useases := useases.UserUsecase(m.s.logger, repository)
	handlers := handlers.UserHandler(useases, m.s.logger)

	return &userModuleModule{
		moduleFactory: m,
		h:             handlers,
	}
}

func (m *userModuleModule) Init() {

	// handlers
	m.r.Get(constants.ROUTE().SALARY, m.h.GetSalary)
	m.r.Post(constants.ROUTE().UserFin, m.h.AddNewSalaryBymonth)

}
