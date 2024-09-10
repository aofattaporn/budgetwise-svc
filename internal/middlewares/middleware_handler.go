package middlewares

import (
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/goproject/pkg/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/goproject/configs"
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/entities"
)

type IMiddlewaresHandler interface {
	Cors() fiber.Handler
	RouterNotFound() fiber.Handler
	Logger() fiber.Handler
	Recover() func(c *fiber.Ctx, e interface{})
}

type middlewaresHandler struct {
	cfg    configs.IConfig
	logger log.ILogger
}

func MiddlewaresHandler(cfg configs.IConfig, logger log.ILogger) IMiddlewaresHandler {
	return &middlewaresHandler{
		cfg:    cfg,
		logger: logger,
	}
}

func (h *middlewaresHandler) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

func (h *middlewaresHandler) RouterNotFound() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(404).JSON(&entities.Response{
			Code: 404,
			Data: nil,
		})
	}
}

func (h *middlewaresHandler) Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// skip if route is healthcheck
		if strings.Contains(c.Path(), constants.ROUTE().HEALTHCHECK) && c.Method() == http.MethodGet {
			return c.Next()
		}

		// startTime := time.Now()
		reqHeaders := make(entities.LogHeaders)
		c.Request().Header.VisitAll(func(key, value []byte) {
			reqHeaders[string(key)] = string(value)
		})

		defer func() {

			// endTime := time.Now()
			// resHeaders := make(entities.LogHeaders)
			// c.Response().Header.VisitAll(func(key, value []byte) {
			// 	resHeaders[string(key)] = string(value)
			// })
			// convert json
			// logJson, _ := json.Marshal(&entities.LoggerRequestAndResponse{
			// 	Method: c.Method(),
			// 	Path:   c.Path(),
			// 	Req: entities.LogReq{
			// 		Headers: reqHeaders,
			// 		Params:  c.Queries(),
			// 		Body:    string(c.Body()),
			// 		Time:    startTime.Format(time.RFC3339Nano),
			// 	},
			// 	Res: entities.LogRes{
			// 		Status:  c.Response().StatusCode(),
			// 		Headers: resHeaders,
			// 		Body:    string(c.Response().Body()),
			// 		Time:    endTime.Format(time.RFC3339Nano),
			// 	},
			// 	LatencyMS: fmt.Sprintf("%d", endTime.Sub(startTime).Milliseconds()),
			// })

			h.logger.Info(string(c.Context().RequestURI()))
		}()

		return c.Next()
	}
}

func (h *middlewaresHandler) Recover() func(c *fiber.Ctx, e interface{}) {
	return func(_ *fiber.Ctx, e interface{}) {
		h.logger.Errorf("recover error : %v ,\n%s", e, debug.Stack())
	}
}
