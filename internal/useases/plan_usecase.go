package useases

import (
	"time"

	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/pkg/log"
)

type IPlanUsecase interface {
	GetAllPlans() entities.PlanList
	CreatePlan(req entities.PlanningRequest) (*entities.PlanList, error)
	UpdatePlan(planId int, req entities.PlanningRequest) (*entities.PlanList, error)
	DeletePlan(deleteId int) (*entities.PlanList, error)
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

func (u *planUsecase) CreatePlan(req entities.PlanningRequest) (*entities.PlanList, error) {
	err := u.r.AddPlan(entities.Plan{
		Name:           req.Name,
		Usage:          0,
		Amount:         req.Amount,
		IconIndex:      req.IconIndex,
		CreateDate:     time.Now(),
		UpdatePlanDate: time.Now(),
		UserID:         1,
		AccountID:      req.AccountID,
	})
	if err != nil {
		u.l.Errorf("create plan error %v", err)
		return nil, err
	}

	plans, err := u.r.FindAllPlans()
	if err != nil {
		u.l.Errorf("find plan error: %v", err)
		return nil, err
	}

	return &plans, nil
}

func (u *planUsecase) UpdatePlan(planId int, req entities.PlanningRequest) (*entities.PlanList, error) {
	err := u.r.UpdatePlan(entities.Plan{
		PlanID:         planId,
		Name:           req.Name,
		Usage:          0,
		Amount:         req.Amount,
		IconIndex:      req.IconIndex,
		UpdatePlanDate: time.Now(),
		UserID:         1,
		AccountID:      req.AccountID,
	})
	if err != nil {
		u.l.Errorf("create plan error %v", err)
		return nil, err
	}

	plans, err := u.r.FindAllPlans()
	if err != nil {
		u.l.Errorf("find plan error: %v", err)
		return nil, err
	}

	return &plans, nil
}

func (u *planUsecase) DeletePlan(deleteId int) (*entities.PlanList, error) {
	err := u.r.DeletePlanById(deleteId)
	if err != nil {
		u.l.Errorf("delete accounts error %v", err)
		return nil, err
	}

	plans, err := u.r.FindAllPlans()
	if err != nil {
		u.l.Errorf("find plan error: %v", err)
		return nil, err
	}

	return &plans, nil
}
