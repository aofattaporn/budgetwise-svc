package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/goproject/internal/customerrors"
	"github.com/goproject/internal/entities"
	"github.com/goproject/internal/useases"
	"github.com/goproject/pkg/log"
)

type IPlanHandler interface {
	GetPlanById(c *fiber.Ctx) error
	GetAllPlans(c *fiber.Ctx) error
	CreatePlan(c *fiber.Ctx) error
	UpdatePlan(c *fiber.Ctx) error
	DeletePlans(c *fiber.Ctx) error
}

type planHandler struct {
	u useases.IPlanUsecase
	l log.ILogger
}

func PlanHandler(usease useases.IPlanUsecase, logger log.ILogger) IPlanHandler {
	return &planHandler{
		u: usease,
		l: logger,
	}
}

// CreatePlan godoc
// @Summary      Create a new plan
// @Description  Create a new planning entry
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        request body entities.PlanningRequest true "Plan information"
// @Success      200  {object}  entities.Response{data=entities.PlanDetails}  "Success response with created plan information"
// @Failure      400  {object}  entities.ErrorResponse  "Invalid input"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /plans [post]
func (h *planHandler) CreatePlan(c *fiber.Ctx) error {
	req := new(entities.PlanningRequest)
	if err := c.BodyParser(req); err != nil {
		return c.JSON(&entities.ErrorResponse{
			Code:         400,
			Timestamp:    time.Now(),
			ErrorMessage: "Invalid request payload",
		})
	}

	plans, err := h.u.CreatePlan(*req)
	if err != nil {
		return c.JSON(&entities.ErrorResponse{
			Code:         1899,
			Timestamp:    time.Now(),
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(&entities.Response{
		Code: 1000,
		Data: plans,
	})
}

// UpdatePlan godoc
// @Summary      Update an existing plan
// @Description  Update a planning entry by ID
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        id path int true "Plan ID"
// @Param        request body entities.PlanningRequest true "Updated plan information"
// @Success      200  {object}  entities.Response{data=entities.PlanDetails}  "Success response with updated plan information"
// @Failure      400  {object}  entities.ErrorResponse  "Invalid input"
// @Failure      404  {object}  entities.ErrorResponse  "Plan not found"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /plans/{id} [put]
func (h *planHandler) UpdatePlan(c *fiber.Ctx) error {
	planId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return customerrors.INVALID_PERAETERS_ERROR("Can't convert plan ID to integer")
	}

	req := new(entities.PlanningRequest)
	if err := c.BodyParser(req); err != nil {
		return c.JSON(&entities.ErrorResponse{
			Code:         400,
			Timestamp:    time.Now(),
			ErrorMessage: "Invalid request payload",
		})
	}

	plans, err := h.u.UpdatePlan(planId, *req)
	if err != nil {
		return c.JSON(&entities.ErrorResponse{
			Code:         1899,
			Timestamp:    time.Now(),
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(&entities.Response{
		Code: 1000,
		Data: plans,
	})
}

// GetPlanById godoc
// @Summary      Get a plan by ID
// @Description  Retrieve a planning entry by its ID
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        id path int true "Plan ID"
// @Success      200  {object}  entities.Response{data=entities.PlanDetails}  "Success response with plan information"
// @Failure      404  {object}  customerrors.CustomError  "Plan not found"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /plans/{id} [get]
func (h *planHandler) GetPlanById(c *fiber.Ctx) error {
	planId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return customerrors.INVALID_PERAETERS_ERROR("Can't convert plan ID to integer")
	}

	plan := h.u.GetPlanById(planId)
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: plan,
	})
}

// GetAllPlans godoc
// @Summary      Get all plans
// @Description  Retrieve all planning entries
// @Tags         plans
// @Accept       json
// @Produce      json
// @Success      200  {object}  entities.Response{data=[]entities.PlanDetails}  "Success response with all plans information"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /plans [get]
func (h *planHandler) GetAllPlans(c *fiber.Ctx) error {
	plans := h.u.GetAllPlans()
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: plans,
	})
}

// DeletePlans godoc
// @Summary      Delete a plan by ID
// @Description  Delete a planning entry by its ID
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        id path int true "Plan ID"
// @Success      200  {object}  entities.Response{data=entities.PlanDetails}  "Success response indicating plan deletion"
// @Failure      404  {object}  customerrors.CustomError  "Plan not found"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /plans/{id} [delete]
func (h *planHandler) DeletePlans(c *fiber.Ctx) error {
	planId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return customerrors.INVALID_PERAETERS_ERROR("Can't convert plan ID to integer")
	}

	plans, err := h.u.DeletePlan(planId)
	if err != nil {
		return c.JSON(&entities.ErrorResponse{
			Code:         1899,
			Timestamp:    time.Now(),
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(&entities.Response{
		Code: 1000,
		Data: plans,
	})
}
