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

// GetAllTransactions godoc
// @Summary      Get all transactions by date
// @Description  Retrieve transactions for a specific date
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        date query string true "Transaction date"
// @Success      200  {object}  entities.Response{data=[]entities.Transaction}  "Success response with list of transactions"
// @Failure      400  {object}  entities.ErrorResponse  "Invalid input parameters"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /transactions [get]
func (h *transactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	date := c.Query("date")
	if date == "" {
		return c.JSON(&entities.ErrorResponse{
			Code:         1799,
			Timestamp:    time.Now(),
			ErrorMessage: "Invalid input parameters",
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

// CreateTransactions godoc
// @Summary      Create a new transaction
// @Description  Create a new transaction with the provided data
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        transaction body entities.TransactionReq true "Transaction Request"
// @Success      200  {object}  entities.Response{data=entities.Transaction}  "Success response with created transaction"
// @Failure      400  {object}  entities.ErrorResponse  "Invalid input"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /transactions [post]
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

// DeleteTransactions godoc
// @Summary      Delete a transaction
// @Description  Delete a transaction by ID
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        id path int true "Transaction ID"
// @Success      200  {object}  entities.Response{data=object}  "Success response indicating deletion"
// @Failure      400  {object}  entities.ErrorResponse  "Invalid ID format"
// @Failure      404  {object}  entities.ErrorResponse  "Transaction not found"
// @Router       /transactions/{id} [delete]
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

// DeleteAllTransactions godoc
// @Summary      Delete all transactions
// @Description  Delete all transactions in the database
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Success      200  {object}  entities.Response{data=object}  "Success response indicating all transactions deletion"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /transactions/all [delete]
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
