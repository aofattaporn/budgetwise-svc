package server

import (
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/handlers"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/internal/useases"
)

type IPlanModule interface {
	Init()
}

type planModule struct {
	*moduleFactory
	h handlers.IPlanHandler
}

func (m *moduleFactory) PlanModule() IPlanModule {

	repository := repositories.PlanRepository(m.s.db.GetDb())
	useases := useases.PlanUsecase(m.s.logger, repository)
	handler := handlers.PlanHandler(useases, m.s.logger)

	return &planModule{
		moduleFactory: m,
		h:             handler,
	}
}

func (m *planModule) Init() {

	// handlers
	m.r.Post(constants.ROUTE().PLANS, m.h.CreatePlan)
	m.r.Get(constants.ROUTE().PLANS, m.h.GetAllPlans)
	m.r.Delete(constants.ROUTE().PLANS_ID, m.h.DeletePlans)

}
