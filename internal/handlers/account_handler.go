package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/useases"
)

type IAccountHandler interface {
	GetAllAccounts(c *fiber.Ctx) error
	CreateAccount(c *fiber.Ctx) error
}

type accountHandler struct {
	u useases.IAccountUsecase
}

func AccountHandler(usease useases.IAccountUsecase) IAccountHandler {
	return &accountHandler{
		u: usease,
	}
}

func (h *accountHandler) GetAllAccounts(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: h.u.GetAllAccounts(),
	})
}

func (h *accountHandler) CreateAccount(c *fiber.Ctx) error {

	req := new(entities.AccountRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	h.u.CreateAccount(*req)

	return c.JSON(&entities.Response{
		Code: 1000,
		Data: nil,
	})
}
