package useases

import (
	"fmt"
	"time"

	"github.com/goproject/internal/customerrors"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/pkg/log"
)

type IUserUsecase interface {
	GetSalaryAndDateReset(userId int, monthYear string) (*entities.UserFinancialsRes, error)
	AddNewSalaryBymonth(req *entities.UserFinancialsReq) (*entities.UserFinancialsRes, error)
}

type userUsecase struct {
	l log.ILogger
	r repositories.IUserRepository
}

func UserUsecase(logger log.ILogger, repository repositories.IUserRepository) IUserUsecase {
	return &userUsecase{
		l: logger,
		r: repository,
	}
}

func (h *userUsecase) GetSalaryAndDateReset(userId int, monthYear string) (*entities.UserFinancialsRes, error) {

	h.l.Infof("get salary and usage from month: %s", monthYear)
	userFin, err := h.r.GetSalaryAndDateReset(userId, monthYear)

	if err != nil {
		h.l.Errorf("get salary and reset date error %v", err)
		return nil, customerrors.TECHNICAL_ERROR("get salary and reset date error")
	}

	if userFin == nil {
		h.l.Errorf("user financials not found in this month")
		return nil, customerrors.DATA_NOT_FOUND("user financials not found in this month")
	}

	return userFin, nil
}

func (u *userUsecase) AddNewSalaryBymonth(req *entities.UserFinancialsReq) (*entities.UserFinancialsRes, error) {

	userId := 1
	u.l.Infof("add salary and usage from month: %s", req.Month)
	err := u.r.AddNewSalaryBymonth(&entities.UserFinancials{
		UserId: userId,
		Salary: req.Salary,
		Month:  time.Date(req.Month.Year(), req.Month.Month(), 1, 0, 0, 0, 0, req.Month.Location()),
		Usages: 0,
	})

	if err != nil {
		u.l.Errorf("add salary and reset date error: %v", err)
		return nil, customerrors.BUSINESS_ERROR("Duplicate month assignment")
	}

	return u.GetSalaryAndDateReset(userId, fmt.Sprintf("%d-%02d", req.Month.Year(), req.Month.Month()))
}
