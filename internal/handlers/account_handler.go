package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/useases"
	"github.com/goproject/pkg/log"
)

type IAccountHandler interface {
	GetAllAccounts(c *fiber.Ctx) error
	CreateAccount(c *fiber.Ctx) error
	UpdateAccount(c *fiber.Ctx) error
	DeleteAccount(c *fiber.Ctx) error
	PatchAccount(c *fiber.Ctx) error
}

type accountHandler struct {
	u useases.IAccountUsecase
	l log.ILogger
}

func AccountHandler(usease useases.IAccountUsecase, logger log.ILogger) IAccountHandler {
	return &accountHandler{
		u: usease,
		l: logger,
	}
}

func (h *accountHandler) GetAllAccounts(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code:        1000,
		Description: constants.STATUS().SUCCESS,
		Data:        h.u.GetAllAccounts(),
	})
}

func (h *accountHandler) CreateAccount(c *fiber.Ctx) error {

	req := new(entities.AccountRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	h.u.CreateAccount(*req)

	return c.JSON(&entities.Response{
		Code:        1000,
		Description: constants.STATUS().SUCCESS,
		Data:        nil,
	})
}

func (h *accountHandler) UpdateAccount(c *fiber.Ctx) error {

	req := new(entities.Account)
	if err := c.BodyParser(req); err != nil {
		return err
	}
	h.u.UpdateAccount(*req)

	return c.JSON(&entities.Response{
		Code:        1000,
		Description: constants.STATUS().SUCCESS,
		Data:        nil,
	})
}

func (h *accountHandler) DeleteAccount(c *fiber.Ctx) error {

	accountId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "convert id error"}
	}
	h.u.DeleteAccount(accountId)

	return c.JSON(&entities.Response{
		Code:        1000,
		Description: constants.STATUS().SUCCESS,
		Data:        nil,
	})
}

func (h *accountHandler) PatchAccount(c *fiber.Ctx) error {

	accountId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "convert id error"}
	}

	req := new(entities.AccountRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	h.u.PatchAccount(accountId, *req)

	return c.JSON(&entities.Response{
		Code:        1000,
		Description: constants.STATUS().SUCCESS,
		Data:        nil,
	})
}
