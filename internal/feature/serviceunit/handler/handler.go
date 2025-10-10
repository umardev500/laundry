package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/serviceunit/contract"
	"github.com/umardev500/laundry/internal/feature/serviceunit/dto"
	"github.com/umardev500/laundry/internal/feature/serviceunit/mapper"
	"github.com/umardev500/laundry/internal/feature/serviceunit/query"
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

// Create handles POST /api/service-units
func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateServiceUnitRequest

	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	unit := req.ToDomain(ctx)

	res, err := h.service.Create(ctx, unit)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToResponse(res))
}

// List handles GET /api/service-units
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListServiceUnitQuery

	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	q.Normalize()
	ctx := appctx.New(c.UserContext())

	data, err := h.service.List(ctx, &q)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		mapper.ToResponseList(data.Data),
		httpx.NewPagination(q.Page, q.Limit, data.Total),
	)
}

// Get handles GET /api/service-units/:id
func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	unit, err := h.service.GetByID(ctx, id)
	if err != nil {
		return httpx.NotFound(c, err.Error())
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(unit))
}

// Update handles PUT /api/service-units/:id
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdateServiceUnitRequest

	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	unit := req.ToDomain(id)

	res, err := h.service.Update(ctx, unit)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

// Delete handles DELETE /api/service-units/:id
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	if err := h.service.Delete(ctx, id); err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.NoContent(c)
}
