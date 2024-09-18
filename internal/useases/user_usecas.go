package useases

import (
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/pkg/log"
)

type IUserUsecase interface {
	GetSalaryAndDateReset(userId int) *entities.SalaryAndResetDate
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

func (h *userUsecase) GetSalaryAndDateReset(userId int) *entities.SalaryAndResetDate {
	user, err := h.r.GetSalaryAndDateReset(userId)
	if err != nil {
		h.l.Errorf("get salary and reset date error %v", err)
	}
	return user
}
