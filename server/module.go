package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/middlewares"
)

type IModuleFactory interface {
	HealthCheckModule()
	UserModule() IUserModule
	AccountModule() IAccountModule
	PlanModule() IPlanModule
	TransactionModule() ITransactionModule
}

type moduleFactory struct {
	r   fiber.Router
	s   *fiberServer
	mid middlewares.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *fiberServer, mid middlewares.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *fiberServer) middlewares.IMiddlewaresHandler {
	return middlewares.MiddlewaresHandler(s.cfg, s.logger)
}
