package useases

import (
	"time"

	"github.com/goproject/internal/customerrors"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/repositories"
	"github.com/goproject/pkg/log"
)

type ITransactionUsecase interface {
	GetTransaction(date string) (entities.TransactionListRes, error)
	CreateTransactions(req entities.TransactionReq) (entities.TransactionListRes, error)
	DeleteTransactions(id int) error
	DeleteAllTransactions() error
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

func (u *transactionsUsecase) GetTransaction(date string) (entities.TransactionListRes, error) {

	u.l.ServiceInfof("get transactions by date %s", date)
	ts, err := u.r.FindListTransaction(date)
	if err != nil {
		u.l.Errorf("find plan error: %s", err)
		return []entities.TransactionRes{}, customerrors.BUSINESS_ERROR(err.Error())
	}

	return ts, nil
}

func (u *transactionsUsecase) CreateTransactions(req entities.TransactionReq) (entities.TransactionListRes, error) {

	u.l.ServiceInfof("create transaction name: (%s)", req.Name)
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
		u.l.Errorf("create transaction error %v", err)
		return entities.TransactionListRes{}, customerrors.TECHNICAL_ERROR(err.Error())
	}

	return entities.TransactionListRes{}, nil
}

func (u *transactionsUsecase) DeleteTransactions(id int) error {
	err := u.r.DeleteTransactionById(id)
	if err != nil {
		u.l.Errorf("delete accounts error %v", err)
		return customerrors.TECHNICAL_ERROR(err.Error())
	}

	return nil
}

func (u *transactionsUsecase) DeleteAllTransactions() error {
	err := u.r.TruncateTransaction()
	if err != nil {
		u.l.Errorf("truncate accounts error %v", err)
		return customerrors.TECHNICAL_ERROR(err.Error())
	}

	return nil
}
