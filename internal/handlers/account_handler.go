package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/entities"
)

type IAccountHandler interface {
	GetAllAccounts(c *fiber.Ctx) error
}

type accountHandler struct {
}

func AccountHandler() IAccountHandler {
	return &accountHandler{}
}

func (h *accountHandler) GetAllAccounts(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: nil,
	})
}
