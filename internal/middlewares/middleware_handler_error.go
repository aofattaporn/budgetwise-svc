package middlewares

import (
	"github.com/goproject/pkg/log"

	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/entities"
)

func MappingError(logger log.ILogger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		if err == nil {
			return c.Next()
		}
		c.Set("content-type", "application/json; charset=utf-8")
		switch e := err.(type) {
		case *fiber.Error:
			logger.Errorf("fiber response error : %+v", e)
			return e
		default:
			logger.Errorf("internal server error : %+v", e)
			internalServerErr := 200
			return c.Status(500).JSON(&entities.Response{
				Code:        internalServerErr,
				Description: "message",
			})
		}
	}
}
