package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/plan/contract"
	"github.com/umardev500/laundry/internal/feature/plan/dto"
	"github.com/umardev500/laundry/internal/feature/plan/mapper"
	"github.com/umardev500/laundry/internal/feature/plan/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"
)

type Handler struct {
	service   contract.Plan
	validator *validator.Validator
}

func NewHandler(s contract.Plan, v *validator.Validator) *Handler {
	return &Handler{
		service:   s,
		validator: v,
	}
}

// Activate handles PATCH /api/plans/:id/activate
func (h *Handler) Activate(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	plan, err := h.service.Activate(ctx, id)
	if err != nil {
		return handlePlanError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(plan))
}

// Deactivate handles PATCH /api/plans/:id/deactivate
func (h *Handler) Deactivate(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	plan, err := h.service.Deactivate(ctx, id)
	if err != nil {
		return handlePlanError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(plan))
}

// Restore handles PATCH /api/plans/:id/restore
func (h *Handler) Restore(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	plan, err := h.service.Restore(ctx, id)
	if err != nil {
		return handlePlanError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(plan))
}

// Create handles POST /api/plans
func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreatePlanRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	plan := req.ToDomain(ctx)

	res, err := h.service.Create(ctx, plan)
	if err != nil {
		return handlePlanError(c, err)
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToResponse(res))
}

// List handles GET /api/plans
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListPlanQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()
	ctx := appctx.New(c.UserContext())

	data, err := h.service.List(ctx, &q)
	if err != nil {
		return handlePlanError(c, err)
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		mapper.ToResponseList(data.Data),
		httpx.NewPagination(q.Page, q.Limit, data.Total),
	)
}

// Get handles GET /api/plans/:id
func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	plan, err := h.service.GetByID(ctx, id)
	if err != nil {
		return handlePlanError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(plan))
}

// Update handles PUT /api/plans/:id
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdatePlanRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	plan := req.ToDomain(ctx, id)

	res, err := h.service.Update(ctx, plan)
	if err != nil {
		return handlePlanError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

// Delete handles DELETE /api/plans/:id
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	if err := h.service.Delete(ctx, id); err != nil {
		return handlePlanError(c, err)
	}
	return httpx.NoContent(c)
}

// Purge handles DELETE /api/plans/:id/purge
func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	if err := h.service.Purge(ctx, id); err != nil {
		return handlePlanError(c, err)
	}
	return httpx.NoContent(c)
}
