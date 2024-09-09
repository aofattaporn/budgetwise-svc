package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/entities"
)

type IPlanHandler interface {
	GetAllPlans(c *fiber.Ctx) error
}

type planHandler struct {
}

func PlanHandler() IPlanHandler {
	return &planHandler{}
}

func (h *planHandler) GetAllPlans(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: nil,
	})
}
