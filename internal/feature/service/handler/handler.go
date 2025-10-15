package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/service/contract"
	"github.com/umardev500/laundry/internal/feature/service/dto"
	"github.com/umardev500/laundry/internal/feature/service/mapper"
	"github.com/umardev500/laundry/internal/feature/service/query"
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

// Create POST /api/services
func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateServiceRequest

	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	d := req.ToDomain(ctx)

	res, err := h.service.Create(ctx, d)
	if err != nil {
		return handleServiceError(c, err)
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToResponse(res))
}

// List GET /api/services
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListServiceQuery

	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	q.Normalize()
	ctx := appctx.New(c.UserContext())

	page, err := h.service.List(ctx, &q)
	if err != nil {
		return handleServiceError(c, err)
	}

	return httpx.JSONPaginated(c, fiber.StatusOK, mapper.ToResponsePage(page).Data, httpx.NewPagination(q.Page, q.Limit, page.Total))
}

// Get GET /api/services/:id
func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	res, err := h.service.GetByID(ctx, id)
	if err != nil {
		return handleServiceError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

// Update PUT /api/services/:id
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdateServiceRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	d := req.ToDomain(id)

	// Convert sentinel price -1 to no change.
	if req.Price == nil {
		d.BasePrice = -1
	}

	res, err := h.service.Update(ctx, d)
	if err != nil {
		return handleServiceError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

// Delete DELETE /api/services/:id (soft delete)
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	if err := h.service.Delete(ctx, id); err != nil {
		return handleServiceError(c, err)
	}

	return httpx.NoContent(c)
}

// Purge DELETE /api/services/:id/purge (hard delete)
func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	if err := h.service.Purge(ctx, id); err != nil {
		return handleServiceError(c, err)
	}

	return httpx.NoContent(c)
}
