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
// @Description  Create a new plan with the provided data
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        plan body entities.PlanningRequest true "Plan Request"
// @Success      200  {object}  entities.Response{data=entities.Plan}  "Success response with created plan"
// @Failure      400  {object}  entities.ErrorResponse  "Invalid input"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /plans [post]
func (h *planHandler) CreatePlan(c *fiber.Ctx) error {
	req := new(entities.PlanningRequest)
	if err := c.BodyParser(req); err != nil {
		return err
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
// @Description  Update a plan with the provided data
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        id   path      int                   true "Plan ID"
// @Param        plan body      entities.PlanningRequest true "Updated Plan Request"
// @Success      200  {object}  entities.Response{data=entities.Plan}  "Success response with updated plan"
// @Failure      400  {object}  entities.ErrorResponse  "Invalid ID format"
// @Failure      500  {object}  entities.ErrorResponse  "Internal server error"
// @Router       /plans/{id} [put]
func (h *planHandler) UpdatePlan(c *fiber.Ctx) error {
	planId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "convert id error"}
	}
	req := new(entities.PlanningRequest)
	if err := c.BodyParser(req); err != nil {
		return err
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
// @Description  Retrieve a plan by its ID
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        id path int true "Plan ID"
// @Success      200  {object}  entities.Response{data=entities.Plan}  "Success response with plan details"
// @Failure      400  {object}  entities.ErrorResponse  "Invalid ID format"
// @Failure      404  {object}  entities.ErrorResponse  "Plan not found"
// @Router       /plans/{id} [get]
func (h *planHandler) GetPlanById(c *fiber.Ctx) error {
	planId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return customerrors.INVALID_PERAETERS_ERROR("can't to find plan id")
	}

	return c.JSON(&entities.Response{
		Code: 1000,
		Data: h.u.GetPlanById(planId),
	})
}

// GetAllPlans godoc
// @Summary      Get all plans
// @Description  Retrieve a list of all plans
// @Tags         plans
// @Accept       json
// @Produce      json
// @Success      200  {object}  entities.Response{data=[]entities.Plan}  "Success response with list of plans"
// @Router       /plans [get]
func (h *planHandler) GetAllPlans(c *fiber.Ctx) error {
	return c.JSON(&entities.Response{
		Code: 1000,
		Data: h.u.GetAllPlans(),
	})
}

// DeletePlans godoc
// @Summary      Delete a plan
// @Description  Delete a plan by ID
// @Tags         plans
// @Accept       json
// @Produce      json
// @Param        id path int true "Plan ID"
// @Success      200  {object}  entities.Response{data=entities.Plan}  "Success response with deleted plan"
// @Failure      400  {object}  entities.ErrorResponse  "Invalid ID format"
// @Failure      404  {object}  entities.ErrorResponse  "Plan not found"
// @Router       /plans/{id} [delete]
func (h *planHandler) DeletePlans(c *fiber.Ctx) error {
	planId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "convert id error"}
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
