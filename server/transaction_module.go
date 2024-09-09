package server

type ITransactionModule interface {
	Init()
}

type transactionModule struct {
	*moduleFactory
	// handler handlers.IHealthHandler
}

func (m *moduleFactory) TransactionModule() ITransactionModule {
	return &planModule{
		moduleFactory: m,
	}
}

func (m *transactionModule) Init() {

	// handler
}
