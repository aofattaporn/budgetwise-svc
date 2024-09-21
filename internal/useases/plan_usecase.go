package useases

import (
	"time"

	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/pkg/log"
)

type IPlanUsecase interface {
	GetAllPlans() entities.PlanList
	CreatePlan(req entities.PlanningRequest) error
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

func (u *planUsecase) GetAllPlans() entities.PlanList {
	plans, err := u.r.FindAllPlans()
	if err != nil {
		u.l.Errorf("find accounts error: %v", err)
	}
	return plans
}

func (u *planUsecase) CreatePlan(req entities.PlanningRequest) error {
	err := u.r.AddPlan(entities.Plan{
		Name:           req.Name,
		Amount:         req.Amount,
		IconIndex:      req.IconIndex,
		CreateDate:     time.Now(),
		UpdatePlanDate: time.Now(),
		UserID:         1,
		AccountID:      req.AccountID,
	})
	if err != nil {
		u.l.Errorf("create accounts error %v", err)
		return err
	}

	return nil
}
