package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/address/contract"
	"github.com/umardev500/laundry/internal/feature/address/domain"
	"github.com/umardev500/laundry/internal/feature/address/dto"
	"github.com/umardev500/laundry/internal/feature/address/mapper"
	"github.com/umardev500/laundry/internal/feature/address/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"
)

type Handler struct {
	service   contract.Address
	validator *validator.Validator
}

func NewHandler(s contract.Address, v *validator.Validator) *Handler {
	return &Handler{
		service:   s,
		validator: v,
	}
}

// Create handles POST /api/addresses
func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateAddressRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	addr, err := req.ToDomain(*ctx.UserID())
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	res, err := h.service.Create(ctx, addr)
	if err != nil {
		return handleAddressError(c, err)
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToResponse(res))
}

// List handles GET /api/addresses
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListAddressQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()
	ctx := appctx.New(c.UserContext())

	data, err := h.service.List(ctx, &q)
	if err != nil {
		return handleAddressError(c, err)
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		mapper.ToResponseList(data.Data),
		httpx.NewPagination(q.Page, q.Limit, data.Total),
	)
}

// Get handles GET /api/addresses/:id
func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	var q query.FindAddressByIDQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	addr, err := h.service.GetByID(ctx, id, &q)
	if err != nil {
		return handleAddressError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(addr))
}

// Update handles PUT /api/addresses/:id
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdateAddressRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	addr := req.ToDomain(id)

	res, err := h.service.Update(ctx, addr)
	if err != nil {
		return handleAddressError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

// Delete handles DELETE /api/addresses/:id
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	if err := h.service.Delete(ctx, id); err != nil {
		return handleAddressError(c, err)
	}
	return httpx.NoContent(c)
}

// Purge handles DELETE /api/addresses/:id/purge
func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	if err := h.service.Purge(ctx, id); err != nil {
		return handleAddressError(c, err)
	}
	return httpx.NoContent(c)
}

// Restore handles PATCH /api/addresses/:id/restore
func (h *Handler) Restore(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}
	ctx := appctx.New(c.UserContext())

	addr, err := h.service.Restore(ctx, id)
	if err != nil {
		return handleAddressError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(addr))
}

// SetPrimary handles PATCH /api/addresses/:id/set-primary
func (h *Handler) SetPrimary(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	addr, err := h.service.SetPrimary(ctx, id, *ctx.UserID())
	if err != nil {
		return handleAddressError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(addr))
}

// GetPrimary handles GET /api/addresses/primary/:user_id
func (h *Handler) GetPrimary(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid user_id")
	}

	ctx := appctx.New(c.UserContext())
	addr, err := h.service.GetPrimaryByUserID(ctx, userID)
	if err != nil {
		return handleAddressError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(addr))
}

// handleAddressError maps domain errors to HTTP responses.
func handleAddressError(c *fiber.Ctx, err error) error {
	return domain.HandleAddressError(c, err)
}
