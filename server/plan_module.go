package server

type IPlanModule interface {
	Init()
}

type planModule struct {
	*moduleFactory
	// handler handlers.IHealthHandler
}

func (m *moduleFactory) PlanModule() IPlanModule {
	return &planModule{
		moduleFactory: m,
	}
}

func (m *planModule) Init() {

	// handler

}
