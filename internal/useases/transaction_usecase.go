package useases

import (
	"errors"
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

func TransactionUsecase(logger log.ILogger, repository repositories.ITransactionRepository, accountRepo repositories.IAccountRepository, planRepo repositories.IPlanRepository) ITransactionUsecase {
	return &transactionsUsecase{
		l:           logger,
		r:           repository,
		accountRepo: accountRepo,
		planRepo:    planRepo,
	}
}

type transactionsUsecase struct {
	l           log.ILogger
	r           repositories.ITransactionRepository
	planRepo    repositories.IPlanRepository
	userRepo    repositories.IUserRepository
	accountRepo repositories.IAccountRepository
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

	u.l.ServiceInfof("update amount plans by id: (%s)", req.PlanId)
	err := u.UpdateAccountPlan(req.PlanId, req.Operation, req.Amount)
	if err != nil {
		u.l.Errorf("[error] update amount plans: %v", err)
		return nil, customerrors.TECHNICAL_ERROR(err.Error())
	}

	u.l.ServiceInfof("update salary account by id: (%s)", req.AccountId)
	err = u.UpdateAccountAmount(req.AccountId, req.Operation, req.Amount)
	if err != nil {
		u.l.Errorf("[error] update salary account: %v", err)
		return nil, customerrors.TECHNICAL_ERROR(err.Error())
	}

	u.l.ServiceInfof("create transaction name: (%s)", req.Name)
	err = u.r.AddTransaction(entities.Transaction{
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

	u.l.ServiceInfof("create transaction success")
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

func (u *transactionsUsecase) UpdateAccountAmount(accountId int, operation string, orignAmmount float64) error {
	account, err := u.accountRepo.GetAccountsById(accountId)
	if err != nil {
		return err
	}

	amount, err := CheckOperations(operation, account.Balance, orignAmmount)
	if err != nil {
		return err
	}

	account.Balance = amount
	err = u.accountRepo.UpdateAccount(account)
	if err != nil {
		return err
	}

	return nil
}

func (u *transactionsUsecase) UpdateAccountPlan(planId int, operation string, orignAmmount float64) error {
	plan, err := u.planRepo.GetPlanById(planId)
	if err != nil {
		return err
	}

	amount, err := CheckOperations(operation, plan.Amount, orignAmmount)
	if err != nil {
		return err
	}

	plan.Amount = amount
	err = u.planRepo.UpdatePlan(plan)
	if err != nil {
		return err
	}

	return nil
}

func CheckOperations(operation string, originalAmount float64, ammount float64) (float64, error) {
	switch operation {
	case "transfer":
		return (originalAmount - ammount), nil
	case "income":
		return (originalAmount + ammount), nil
	case "change":
		// TODO: handler on change
		return originalAmount, nil
	default:
		return 0, errors.New("can't to mapping operation error")
	}
}
