package server

type IUserModule interface {
	Init()
}

type userModuleModule struct {
	*moduleFactory
	// handler handlers.IHealthHandler
}

func (m *moduleFactory) UserModule() IUserModule {
	return &userModuleModule{
		moduleFactory: m,
	}
}

func (m *userModuleModule) Init() {

	// handler

}
