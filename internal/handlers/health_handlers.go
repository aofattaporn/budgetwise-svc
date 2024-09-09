package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goproject/configs"
	"github.com/goproject/internal/entities"
)

type IHealthHandler interface {
	HeathCheckHandler(c *fiber.Ctx) error
}

type healthHandler struct {
	cfg configs.IAppConfig
}

func HealthCheckHandler(cfg configs.IAppConfig) IHealthHandler {
	return &healthHandler{
		cfg: cfg,
	}
}

func (h *healthHandler) HeathCheckHandler(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: &entities.HealthResponse{
			Name:    h.cfg.Name(),
			Version: h.cfg.Version(),
		},
	})
}
