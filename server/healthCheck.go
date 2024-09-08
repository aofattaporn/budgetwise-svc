package server

import (
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/handlers"
)

func (m *moduleFactory) HealthCheckModule() {
	h := handlers.HealthCheckHandler(m.s.cfg.App())

	m.r.Get(constants.ROUTE().HEALTHCHECK, h.HeathCheckHandler)
}
