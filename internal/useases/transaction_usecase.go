package useases

import (
	"time"

	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/pkg/log"
)

type ITransactionUsecase interface {
	GetTransaction() entities.TransactionListRes
	CreateTransactions(req entities.TransactionReq) (*entities.TransactionListRes, error)
}

func TransactionUsecase(logger log.ILogger, repository repositories.ITransactionRepository) ITransactionUsecase {
	return &transactionsUsecase{
		l: logger,
		r: repository,
	}
}

type transactionsUsecase struct {
	l log.ILogger
	r repositories.ITransactionRepository
}

func (u *transactionsUsecase) GetTransaction() entities.TransactionListRes {

	ts, err := u.r.FindListTransaction()
	if err != nil {
		u.l.Errorf("find plan error: %v", err)
	}

	return ts
}

func (u *transactionsUsecase) CreateTransactions(req entities.TransactionReq) (*entities.TransactionListRes, error) {

	err := u.r.AddTransaction(entities.Transaction{
		Name:       req.Name,
		Amount:     req.Amount,
		Operation:  req.Operation,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
		UserID:     1,
		PlanId:     req.PlanId,
		AccountId:  req.AccountId,
	})
	if err != nil {
		u.l.Errorf("create plan error %v", err)
		return nil, err
	}

	return nil, nil
}
