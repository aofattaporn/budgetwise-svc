package server

import (
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/handlers"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/internal/useases"
)

type ITransactionModule interface {
	Init()
}

type transactionModule struct {
	*moduleFactory
	h handlers.ITransactionHandler
}

func (m *moduleFactory) TransactionModule() ITransactionModule {

	repository := repositories.TransactionRepository(m.s.db.GetDb())
	acRepo := repositories.AccountRepository(m.s.db.GetDb())
	planRepo := repositories.PlanRepository(m.s.db.GetDb())

	useases := useases.TransactionUsecase(m.s.logger, repository, acRepo, planRepo)
	handler := handlers.TransactionHandler(useases, m.s.logger)

	return &transactionModule{
		moduleFactory: m,
		h:             handler,
	}
}

func (m *transactionModule) Init() {

	// handler
	m.r.Get(constants.ROUTE().TRANSACTIONS, m.h.GetAllTransactions)
	m.r.Post(constants.ROUTE().TRANSACTIONS, m.h.CreateTransactions)
	m.r.Delete(constants.ROUTE().TRANSACTIONS, m.h.DeleteTransactions)
	m.r.Delete(constants.ROUTE().TRANSACTIONS_ID, m.h.DeleteTransactions)

}
