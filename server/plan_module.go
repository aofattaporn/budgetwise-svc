package server

import (
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/handlers"
)

type IPlanModule interface {
	Init()
}

type planModule struct {
	*moduleFactory
	h handlers.IPlanHandler
	// handler handlers.IHealthHandler
}

func (m *moduleFactory) PlanModule() IPlanModule {

	handler := handlers.PlanHandler()

	return &planModule{
		moduleFactory: m,
		h:             handler,
	}
}

func (m *planModule) Init() {

	// handlers
	m.r.Get(constants.ROUTE().PLANS, m.h.GetAllPlans)

}
