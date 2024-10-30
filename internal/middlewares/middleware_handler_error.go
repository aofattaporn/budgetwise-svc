package middlewares

import (
	"time"

	"github.com/goproject/pkg/log"

	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/customerrors"
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

		// ********** CUSTOM ERROR ***********
		case *customerrors.CustomError:
			switch e.ErrorType {

			// *********** business error ***********
			case customerrors.ERROR_TYPE().INVALID_PARAMETER_ERROR:
				logger.Errorf("fiber response error : %+v", e)
				return c.Status(400).JSON(&entities.ErrorResponse{
					Code:         1799,
					Timestamp:    time.Now(),
					ErrorMessage: e.Description,
				})

			// *********** business error ***********
			case customerrors.ERROR_TYPE().BUSINESS_ERROR:
				logger.Errorf("[error] business error : %+v", e)
				return c.Status(400).JSON(&entities.ErrorResponse{
					Code:         1899,
					Timestamp:    time.Now(),
					ErrorMessage: e.Description,
				})

			// *********** technical error ***********
			case customerrors.ERROR_TYPE().Technical_ERROR:
				logger.Errorf("[error] technical error : %+v", e)
				return c.Status(500).JSON(&entities.ErrorResponse{
					Code:         1999,
					Timestamp:    time.Now(),
					ErrorMessage: e.Description,
				})

			default:
				logger.Errorf("internal server error : %+v", e)
				return c.Status(500).JSON(&entities.ErrorResponse{
					Code:         e.Code,
					Timestamp:    time.Now(),
					ErrorMessage: e.Description,
				})
			}

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
