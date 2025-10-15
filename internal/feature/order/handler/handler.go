package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/contract"
	"github.com/umardev500/laundry/internal/feature/order/dto"
	"github.com/umardev500/laundry/internal/feature/order/mapper"
	"github.com/umardev500/laundry/internal/feature/order/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/utils"
	"github.com/umardev500/laundry/pkg/validator"

	historyContract "github.com/umardev500/laundry/internal/feature/orderstatushistory/contract"
	historyMapper "github.com/umardev500/laundry/internal/feature/orderstatushistory/mapper"
	historyQuery "github.com/umardev500/laundry/internal/feature/orderstatushistory/query"
)

type Handler struct {
	service        contract.OrderService
	historyService historyContract.StatusHistoryService
	validator      *validator.Validator
}

func NewHandler(
	service contract.OrderService,
	historyService historyContract.StatusHistoryService,
	validator *validator.Validator,
) *Handler {
	return &Handler{
		service:        service,
		validator:      validator,
		historyService: historyService,
	}
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
		return handleOrderError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(result))
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
		return handleOrderError(c, err)
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToResponse(result))
}

func (h *Handler) History(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid order ID")
	}

	var q historyQuery.OrderStatusHistoryListQuery

	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	q.OrderID = utils.NilIfUUIDZero(id)
	q.Normalize()

	ctx := appctx.New(c.UserContext())

	page, err := h.historyService.List(ctx, &q)
	if err != nil {
		return handleOrderError(c, err)
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		historyMapper.ToResponsePage(page, h.refToResponse).Data,
		httpx.NewPagination(q.Page, q.Limit, page.Total),
	)
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
		return handleOrderError(c, err)
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
		return handleOrderError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(result))
}

// UpdateStatus PATCH /api/orders/:id/status
func (h *Handler) UpdateStatus(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid order ID")
	}

	var q query.UpdateStatusQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	m, err := q.ToDomain(id)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	res, err := h.service.UpdateStatus(ctx, m)
	if err != nil {
		return handleOrderError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

// -----------------------
// Helper methods
// -----------------------

func (h *Handler) refToResponse(ref any) any {
	switch r := ref.(type) {
	case *ent.Order:
		order := mapper.FromEnt(r)
		return mapper.ToResponse(order)
	}
	return ref
}
