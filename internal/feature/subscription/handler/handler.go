package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/subscription/contract"
	"github.com/umardev500/laundry/internal/feature/subscription/dto"
	"github.com/umardev500/laundry/internal/feature/subscription/mapper"
	"github.com/umardev500/laundry/internal/feature/subscription/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"
)

type Handler struct {
	service   contract.Service
	validator *validator.Validator
}

func NewHandler(s contract.Service, v *validator.Validator) *Handler {
	return &Handler{
		service:   s,
		validator: v,
	}
}

// Create handles POST /api/subscriptions
func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateSubscriptionRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	sub := req.ToDomain(ctx)

	created, err := h.service.Create(ctx, sub)
	if err != nil {
		return handleSubscriptionError(c, err)
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToResponse(created))
}

// Get handles GET /api/subscriptions/:id
func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var q query.FindSubscriptionByIDQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	ctx := appctx.New(c.UserContext())

	sub, err := h.service.FindByID(ctx, id, &q)
	if err != nil {
		return handleSubscriptionError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(sub))
}

// List handles GET /api/subscriptions
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListSubscriptionQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()
	ctx := appctx.New(c.UserContext())

	data, err := h.service.List(ctx, &q)
	if err != nil {
		return handleSubscriptionError(c, err)
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		mapper.ToResponseList(data.Data),
		httpx.NewPagination(q.Page, q.Limit, data.Total),
	)
}

// Update handles PUT /api/subscriptions/:id
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdateSubscriptionRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	sub := req.ToDomain(ctx, id)

	updated, err := h.service.Update(ctx, sub)
	if err != nil {
		return handleSubscriptionError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(updated))
}

// UpdateStatus handles PATCH /api/subscriptions/:id/status/:status
func (h *Handler) UpdateStatus(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var q query.UpdateStatusQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	sub, err := q.ToDomain(id)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	updated, err := h.service.UpdateStatus(ctx, sub)
	if err != nil {
		return handleSubscriptionError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(updated))
}

// Delete handles DELETE /api/subscriptions/:id (soft delete)
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	if err := h.service.Delete(ctx, id); err != nil {
		return handleSubscriptionError(c, err)
	}

	return httpx.NoContent(c)
}

// Restore handles PATCH /api/subscriptions/:id/restore
func (h *Handler) Restore(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	data, err := h.service.Restore(ctx, id)
	if err != nil {
		return handleSubscriptionError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(data))
}

// Purge handles DELETE /api/subscriptions/:id/purge (hard delete)
func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	if err := h.service.Purge(ctx, id); err != nil {
		return handleSubscriptionError(c, err)
	}

	return httpx.NoContent(c)
}
