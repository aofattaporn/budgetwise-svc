package server

import (
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/handlers"
)

type ITransactionModule interface {
	Init()
}

type transactionModule struct {
	*moduleFactory
	h handlers.ITransactionHandler
}

func (m *moduleFactory) TransactionModule() ITransactionModule {
	handler := handlers.TransactionHandler()

	return &transactionModule{
		moduleFactory: m,
		h:             handler,
	}
}

func (m *transactionModule) Init() {

	// handler
	m.r.Get(constants.ROUTE().TRANSACTIONS, m.h.GetAllTransactions)
}
