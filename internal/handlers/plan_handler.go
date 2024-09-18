package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/useases"
	"github.com/goproject/pkg/log"
)

type IPlanHandler interface {
	GetAllPlans(c *fiber.Ctx) error
	CreatePlan(c *fiber.Ctx) error
}

type planHandler struct {
	u useases.IPlanUsecase
	l log.ILogger
}

func PlanHandler(usease useases.IPlanUsecase, logger log.ILogger) IPlanHandler {
	return &planHandler{
		u: usease,
		l: logger,
	}
}

func (h *planHandler) CreatePlan(c *fiber.Ctx) error {

	req := new(entities.PlanningRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	h.u.CreatePlan(*req)

	return c.JSON(&entities.Response{
		Code: 1000,
		Data: nil,
	})
}

func (h *planHandler) GetAllPlans(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: nil,
	})
}
