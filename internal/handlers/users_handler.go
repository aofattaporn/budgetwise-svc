package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/entities"
)

type IUserHandler interface {
	GetSalary(c *fiber.Ctx) error
}

type userHandler struct {
}

func UserHandler() IUserHandler {
	return &userHandler{}
}

func (h *userHandler) GetSalary(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: nil,
	})
}
