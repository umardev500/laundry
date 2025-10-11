package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/contract"
	"github.com/umardev500/laundry/internal/feature/order/dto"
	"github.com/umardev500/laundry/internal/feature/order/mapper"
	"github.com/umardev500/laundry/internal/feature/order/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"
)

type Handler struct {
	service   contract.OrderService
	validator *validator.Validator
}

func NewHandler(service contract.OrderService, validator *validator.Validator) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) GuestOrder(c *fiber.Ctx) error {
	var req dto.CreateGuestOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	return c.JSON(req)
}

// List handles GET /orders requests with filtering, sorting, and pagination
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListOrderQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	q.Normalize() // set defaults if not provided

	ctx := appctx.New(c.UserContext())
	result, err := h.service.List(ctx, &q)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	// Convert domain orders to DTOs
	dtoPage := mapper.ToResponsePage(result)

	return httpx.JSONPaginated(c, int(fiber.StatusOK), dtoPage.Data, httpx.NewPagination(q.Page, q.Limit, result.Total))
}
