package useases

import (
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/pkg/log"
)

type IPlanUsecase interface {
	GetAllPlans() entities.Plan
	CreatePlan(req entities.PlanningRequest)
}

func PlanUsecase(logger log.ILogger, repository repositories.IPlanRepository) IPlanUsecase {
	return &planUsecase{
		l: logger,
		r: repository,
	}
}

type planUsecase struct {
	l log.ILogger
	r repositories.IPlanRepository
}

func (u *planUsecase) GetAllPlans() entities.Plan {
	return entities.Plan{}
}

func (u *planUsecase) CreatePlan(req entities.PlanningRequest) {
	err := u.r.AddPlan(entities.Plan{
		Name:      req.Name,
		Amount:    req.Amount,
		UserID:    0,
		AccountID: &req.AccountID,
	})
	if err != nil {
		u.l.Errorf("create accounts error %v", err)
	}
}
