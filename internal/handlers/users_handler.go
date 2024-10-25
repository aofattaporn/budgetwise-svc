package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/constants"
	"github.com/goproject/internal/customerrors"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/useases"
	"github.com/goproject/pkg/log"
)

type IUserHandler interface {
	GetSalary(c *fiber.Ctx) error
}

type userHandler struct {
	u useases.IUserUsecase
	l log.ILogger
}

func UserHandler(usease useases.IUserUsecase, logger log.ILogger) IUserHandler {
	return &userHandler{
		u: usease,
		l: logger,
	}
}

// GetSalary godoc
// @Summary      Get user salary information
// @Description  Retrieve the salary details for a specific user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  entities.Response{data=entities.SalaryAndResetDate}  "Success response with user's salary information"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /users/salary [get]
func (h *userHandler) GetSalary(c *fiber.Ctx) error {

	monthYear := c.Query("monthYear")
	if monthYear == "" {
		return customerrors.INVALID_PERAETERS_ERROR("Invalid input parameters")
	}

	userInfo, err := h.u.GetSalaryAndDateReset(1, monthYear)
	if err != nil {
		return err
	}

	return c.JSON(&entities.Response{
		Code:        1000,
		Description: constants.STATUS().SUCCESS,
		Data:        userInfo,
	})
}
