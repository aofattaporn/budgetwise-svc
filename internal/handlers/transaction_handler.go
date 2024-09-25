package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/useases"
	"github.com/goproject/pkg/log"
)

type ITransactionHandler interface {
	CreateTransactions(c *fiber.Ctx) error
	GetAllTransactions(c *fiber.Ctx) error
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
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: h.u.GetTransaction(),
	})
}

func (h *transactionHandler) CreateTransactions(c *fiber.Ctx) error {

	req := new(entities.TransactionReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	ts, err := h.u.CreateTransactions(*req)
	if err != nil {
		return c.JSON(&entities.ErrorResponse{
			Code:         1899,
			Timestamp:    time.Now(),
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(&entities.Response{
		Code: 1000,
		Data: ts,
	})
}
