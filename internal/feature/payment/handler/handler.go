package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/payment/contract"
	"github.com/umardev500/laundry/internal/feature/payment/domain"
	"github.com/umardev500/laundry/internal/feature/payment/mapper"
	"github.com/umardev500/laundry/internal/feature/payment/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"

	orderMapper "github.com/umardev500/laundry/internal/feature/order/mapper"
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

func (h *Handler) FindById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid payment ID")
	}

	var q query.FindPaymentByIdQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())

	p, err := h.service.GetByID(ctx, id, &q)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPaymentNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrPaymentDeleted):
			return httpx.Forbidden(c, err.Error())
		}

		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(p, h.refToResponse))
}

// List GET /api/payments
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListPaymentQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	q.Normalize()
	ctx := appctx.New(c.UserContext())

	page, err := h.service.List(ctx, &q)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		mapper.ToResponsePage(page, h.refToResponse).Data,
		httpx.NewPagination(q.Page, q.Limit, page.Total),
	)
}

// -----------------------
// Helper methods
// -----------------------

func (h *Handler) refToResponse(ref any) any {
	switch r := ref.(type) {
	case *ent.Order:
		order := orderMapper.FromEnt(r)
		return orderMapper.ToResponse(order)
	}
	return ref
}
