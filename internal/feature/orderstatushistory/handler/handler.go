package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/contract"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/domain"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/mapper"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/query"
	"github.com/umardev500/laundry/pkg/httpx"

	orderMapper "github.com/umardev500/laundry/internal/feature/order/mapper"
)

type Handler struct {
	service contract.StatusHistoryService
}

func NewHandler(s contract.StatusHistoryService) *Handler {
	return &Handler{
		service: s,
	}
}

// GetByID GET /api/order-status-history/:id
func (h *Handler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var q query.StatusHistoryByIDQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())

	res, err := h.service.FindByID(ctx, id, &q)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrStatusHistoryNotFound):
			return httpx.NotFound(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(
		c,
		fiber.StatusOK,
		mapper.ToResponse(res, h.refToResponse),
	)
}

// List GET /api/order-status-history
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.OrderStatusHistoryListQuery

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
