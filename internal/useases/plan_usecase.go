package useases

import (
	"time"

	"github.com/goproject/internal/customerrors"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/pkg/log"
)

type IPlanUsecase interface {
	GetPlanById(planId int) entities.Plan
	GetAllPlans(monthYear string) (*entities.PlanList, error)
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

func (u *planUsecase) GetPlanById(planId int) entities.Plan {
	plans, err := u.r.GetPlanById(planId)
	if err != nil {
		u.l.Errorf("find plan error: %v", err)
	}
	return plans
}

func (u *planUsecase) GetAllPlans(monthYear string) (*entities.PlanList, error) {
	plans, err := u.r.FindAllPlans(monthYear)
	if err != nil {
		u.l.Errorf("repository find all plans: %v", err)
		return nil, customerrors.BUSINESS_ERROR(err.Error())
	}

	if plans == nil {
		u.l.Infof("no plans found")
		return nil, customerrors.DATA_NOT_FOUND("no plans found in this month")
	}

	return &plans, nil
}

func (u *planUsecase) CreatePlan(req entities.PlanningRequest) (*entities.PlanList, error) {
	err := u.r.AddPlan(entities.Plan{
		Name:       req.Name,
		Type:       req.Type,
		Usage:      0,
		Amount:     req.Amount,
		IconIndex:  req.IconIndex,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
		UserID:     1,
		Month:      req.Month,
		AccountID:  req.AccountID,
	})

	if err != nil {
		u.l.Errorf("create plan error %v", err)
		return nil, err
	}

	formattedDate := time.Now().Format("2006-01")
	plans, err := u.r.FindAllPlans(formattedDate)
	if err != nil {
		u.l.Errorf("find plan error: %v", err)
		return nil, err
	}

	return &plans, nil
}

func (u *planUsecase) UpdatePlan(planId int, req entities.PlanningRequest) (*entities.PlanList, error) {
	err := u.r.UpdatePlan(entities.Plan{
		Id:         planId,
		Name:       req.Name,
		Usage:      0,
		Amount:     req.Amount,
		IconIndex:  req.IconIndex,
		UpdateDate: time.Now(),
		UserID:     1,
		AccountID:  req.AccountID,
	})
	if err != nil {
		u.l.Errorf("create plan error %v", err)
		return nil, err
	}

	plans, err := u.r.FindAllPlans("2024-10")
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

	plans, err := u.r.FindAllPlans("2024-10")
	if err != nil {
		u.l.Errorf("find plan error: %v", err)
		return nil, err
	}

	return &plans, nil
}
