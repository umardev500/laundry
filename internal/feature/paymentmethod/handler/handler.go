package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/contract"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/domain"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/dto"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/mapper"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/query"
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

// Create POST /api/payment-methods
func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreatePaymentMethodRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	pm := req.ToDomain()

	res, err := h.service.Create(ctx, pm)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPaymentMethodAlreadyExists):
			return httpx.Conflict(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToResponse(res))
}

// List GET /api/payment-methods
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListPaymentMethodQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	q.Normalize()
	ctx := appctx.New(c.UserContext())

	page, err := h.service.List(ctx, &q)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSONPaginated(c, fiber.StatusOK, mapper.ToResponseList(page.Data),
		httpx.NewPagination(q.Page, q.Limit, page.Total))
}

// Get GET /api/payment-methods/:id
func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	pm, err := h.service.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPaymentMethodNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrPaymentMethodDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(pm))
}

// Update PUT /api/payment-methods/:id
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdatePaymentMethodRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	pm := req.ToDomain(id)

	res, err := h.service.Update(ctx, pm)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPaymentMethodNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrPaymentMethodDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

// Delete DELETE /api/payment-methods/:id (soft delete)
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	if err := h.service.Delete(ctx, id); err != nil {
		switch {
		case errors.Is(err, domain.ErrPaymentMethodNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrPaymentMethodDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.NoContent(c)
}

// Purge DELETE /api/payment-methods/:id/purge (hard delete)
func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	if err := h.service.Purge(ctx, id); err != nil {
		switch {
		case errors.Is(err, domain.ErrPaymentMethodNotFound):
			return httpx.NotFound(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.NoContent(c)
}
