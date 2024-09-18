package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/useases"
	"github.com/goproject/pkg/log"
)

type IUserHandler interface {
	GetSalary(c *fiber.Ctx) error
}

type userHandler struct {
	u useases.IUserUsecase
	l log.ILogger
}

func UserHandler(usease useases.IUserUsecase, logger log.ILogger) IUserHandler {
	return &userHandler{
		u: usease,
		l: logger,
	}
}

func (h *userHandler) GetSalary(c *fiber.Ctx) error {
	user := h.u.GetSalaryAndDateReset(1)
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: &user,
	})
}
