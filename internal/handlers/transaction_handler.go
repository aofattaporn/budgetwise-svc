package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/entities"
)

type ITransactionHandler interface {
	GetAllTransactions(c *fiber.Ctx) error
}

type transactionHandler struct {
}

func TransactionHandler() ITransactionHandler {
	return &transactionHandler{}
}

func (h *transactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: nil,
	})
}
