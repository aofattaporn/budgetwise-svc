package server

type IAccountModule interface {
	Init()
}

type accountModule struct {
	*moduleFactory
	// handler handlers.IHealthHandler
}

func (m *moduleFactory) AccountModule() IAccountModule {
	return &accountModule{
		moduleFactory: m,
	}
}

func (m *accountModule) Init() {

	// handler

}
