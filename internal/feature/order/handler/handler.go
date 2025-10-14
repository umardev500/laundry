package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/contract"
	"github.com/umardev500/laundry/internal/feature/order/domain"
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

	// Validate
	if err := req.Validate(); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	data, err := req.ToDomain(uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"))
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	result, err := h.service.GuestOrder(ctx, data)
	if err != nil {
		switch e := err.(type) {
		case *domain.ServiceUnavailableError:
			return httpx.BadRequest(c, e.Error())
		}

		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToResponse(result))
}

func (h *Handler) FindByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return httpx.BadRequest(c, "invalid order ID")
	}

	var q query.OrderQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	result, err := h.service.FindByID(ctx, id, &q)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrOrderNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrOrderDeleted):
			return httpx.BadRequest(c, err.Error())
		case errors.Is(err, domain.ErrUnauthorizedOrderAccess):
			return httpx.Forbidden(c, err.Error())
		}

		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(result))
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

func (h *Handler) Preview(c *fiber.Ctx) error {
	var req dto.PreviewOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	// Validate
	if err := req.Validate(); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	data, err := req.ToDomain()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	result, err := h.service.Preview(ctx, data)
	if err != nil {
		switch e := err.(type) {
		case *domain.ServiceUnavailableError:
			return httpx.BadRequest(c, e.Error())
		}

		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(result))
}
