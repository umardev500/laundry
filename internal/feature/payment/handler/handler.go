package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/payment/contract"
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
		mapper.ToResponsePage(page, func(a any) any {
			switch ref := a.(type) {
			case *ent.Order:
				order := orderMapper.FromEnt(ref)
				return orderMapper.ToResponse(order)
			}

			return a
		}).Data,
		httpx.NewPagination(q.Page, q.Limit, page.Total),
	)
}
