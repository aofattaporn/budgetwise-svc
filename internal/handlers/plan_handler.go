package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/useases"
	"github.com/goproject/pkg/log"
)

type IPlanHandler interface {
	GetAllPlans(c *fiber.Ctx) error
	CreatePlan(c *fiber.Ctx) error
	UpdatePlan(c *fiber.Ctx) error
	DeletePlans(c *fiber.Ctx) error
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

	plans, err := h.u.CreatePlan(*req)
	if err != nil {
		return c.JSON(&entities.ErrorResponse{
			Code:         1899,
			Timestamp:    time.Now(),
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(&entities.Response{
		Code: 1000,
		Data: plans,
	})
}

func (h *planHandler) UpdatePlan(c *fiber.Ctx) error {

	planId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "convert id error"}
	}
	req := new(entities.PlanningRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	plans, err := h.u.UpdatePlan(planId, *req)
	if err != nil {
		return c.JSON(&entities.ErrorResponse{
			Code:         1899,
			Timestamp:    time.Now(),
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(&entities.Response{
		Code: 1000,
		Data: plans,
	})
}

func (h *planHandler) GetAllPlans(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: h.u.GetAllPlans(),
	})
}

func (h *planHandler) DeletePlans(c *fiber.Ctx) error {

	planId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "convert id error"}
	}

	plans, err := h.u.DeletePlan(planId)
	if err != nil {
		return c.JSON(&entities.ErrorResponse{
			Code:         1899,
			Timestamp:    time.Now(),
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(&entities.Response{
		Code: 1000,
		Data: plans,
	})
}
