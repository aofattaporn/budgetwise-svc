package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/customerrors"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/useases"
	"github.com/goproject/pkg/log"
)

type ITransactionHandler interface {
	CreateTransactions(c *fiber.Ctx) error
	GetAllTransactions(c *fiber.Ctx) error
	DeleteTransactions(c *fiber.Ctx) error
	DeleteAllTransactions(c *fiber.Ctx) error
}

type transactionHandler struct {
	u useases.ITransactionUsecase
	l log.ILogger
}

func TransactionHandler(useasecase useases.ITransactionUsecase, logger log.ILogger) ITransactionHandler {
	return &transactionHandler{
		u: useasecase,
		l: logger,
	}
}

func (h *transactionHandler) GetAllTransactions(c *fiber.Ctx) error {

	date := c.Query("date")
	if date == "" {
		return c.JSON(&entities.ErrorResponse{
			Code:         1799,
			Timestamp:    time.Now(),
			ErrorMessage: "Invalid input parammeters",
		})
	}

	ts, err := h.u.GetTransaction(date)
	if err != nil {
		return err
	}

	return c.JSON(&entities.Response{
		Code:        1000,
		Description: "Success",
		Data:        ts,
	})
}

func (h *transactionHandler) CreateTransactions(c *fiber.Ctx) error {

	req := new(entities.TransactionReq)
	if err := c.BodyParser(req); err != nil {
		return customerrors.INVALID_PERAETERS_ERROR(err.Error())
	}

	ts, err := h.u.CreateTransactions(*req)
	if err != nil {
		return err
	}

	return c.JSON(&entities.Response{
		Code:        1000,
		Description: "Success",
		Data:        ts,
	})
}

func (h *transactionHandler) DeleteTransactions(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return customerrors.INVALID_PERAETERS_ERROR(err.Error())
	}

	err = h.u.DeleteTransactions(id)
	if err != nil {
		return err
	}

	return c.JSON(&entities.Response{
		Code:        1000,
		Description: "Success",
		Data:        nil,
	})
}

func (h *transactionHandler) DeleteAllTransactions(c *fiber.Ctx) error {

	err := h.u.DeleteAllTransactions()
	if err != nil {
		return err
	}

	return c.JSON(&entities.Response{
		Code:        1000,
		Description: "Success",
		Data:        nil,
	})
}
