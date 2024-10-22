package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/customerrors"
	"github.com/goproject/internal/entities" // Import where `Response` is defined
	"github.com/goproject/internal/useases"
	"github.com/goproject/pkg/log"
)

type IAccountHandler interface {
	GetAllAccounts(c *fiber.Ctx) error
	CreateAccount(c *fiber.Ctx) error
	DeleteAllAccounts(c *fiber.Ctx) error
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

// ShowAccount godoc
// @Summary      Show all accounts
// @Description  Get all accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  entities.Response{data=[]entities.Account}
// @Router       /accounts [get]
func (h *accountHandler) GetAllAccounts(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code:        1000,
		Description: constants.STATUS().SUCCESS,
		Data:        h.u.GetAllAccounts(), // AccountsList
	})
}

// CreateAccount godoc
// @Summary      Create a new account
// @Description  Create an account with the provided data
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account body entities.AccountRequest true "Account Request"
// @Success      200  {object}  entities.Response{data=object}
// @Router       /accounts [post]
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

// DeleteAccount godoc
// @Summary      Delete an account
// @Description  Delete an account by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id path int true "Account ID"
// @Success      200  {object}  entities.Response{data=object}
// @Router       /accounts/{id} [delete]
func (h *accountHandler) DeleteAccount(c *fiber.Ctx) error {
	accountId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "convert id error"}
	}

	err = h.u.DeleteAccount(accountId)
	if err != nil {
		return customerrors.BUSINESS_ERROR(err.Error())
	}

	return c.JSON(&entities.Response{
		Code:        1000,
		Description: constants.STATUS().SUCCESS,
	})
}

// DeleteAllAccounts godoc
// @Summary      Delete all accounts
// @Description  Delete all accounts from the system
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  entities.Response{data=object}
// @Router       /accounts [delete]
func (h *accountHandler) DeleteAllAccounts(c *fiber.Ctx) error {
	h.u.DeleteAllAccounts()
	return c.JSON(&entities.Response{
		Code:        1000,
		Description: constants.STATUS().SUCCESS,
		Data:        nil,
	})
}

// PatchAccount godoc
// @Summary      Partially update an account
// @Description  Update specific fields of an account by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id path int true "Account ID"
// @Param        account body entities.AccountRequest true "Partial Account Request"
// @Success      200  {object}  entities.Response{data=object}
// @Router       /accounts/{id} [patch]
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
