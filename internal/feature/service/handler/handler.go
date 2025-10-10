package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/service/contract"
	"github.com/umardev500/laundry/internal/feature/service/domain"
	"github.com/umardev500/laundry/internal/feature/service/dto"
	"github.com/umardev500/laundry/internal/feature/service/mapper"
	"github.com/umardev500/laundry/internal/feature/service/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"

	queryPkg "github.com/umardev500/laundry/pkg/query"
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
		switch {
		case errors.Is(err, domain.ErrServiceAlreadyExists):
			return httpx.Conflict(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
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
		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSONPaginated(c, fiber.StatusOK, mapper.ToResponsePage(page).Data, httpx.NewPagination(q.Page, q.Limit, page.Total))
}

// Get GET /api/services/:id
func (h *Handler) Get(c *fiber.Ctx) error {
	var idQuery queryPkg.GetByIDQuery
	if err := c.ParamsParser(&idQuery); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	id, err := idQuery.UUID()
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	res, err := h.service.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrServiceNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrServiceDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

// Update PUT /api/services/:id
func (h *Handler) Update(c *fiber.Ctx) error {
	var idQuery queryPkg.GetByIDQuery
	if err := c.ParamsParser(&idQuery); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	id, err := idQuery.UUID()
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
		switch {
		case errors.Is(err, domain.ErrServiceNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrServiceDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

// Delete DELETE /api/services/:id (soft delete)
func (h *Handler) Delete(c *fiber.Ctx) error {
	var idQuery queryPkg.GetByIDQuery
	if err := c.ParamsParser(&idQuery); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	id, err := idQuery.UUID()
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	if err := h.service.Delete(ctx, id); err != nil {
		switch {
		case errors.Is(err, domain.ErrServiceNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrServiceDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.NoContent(c)
}

// Purge DELETE /api/services/:id/purge (hard delete)
func (h *Handler) Purge(c *fiber.Ctx) error {
	var idQuery queryPkg.GetByIDQuery
	if err := c.ParamsParser(&idQuery); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	id, err := idQuery.UUID()
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	if err := h.service.Purge(ctx, id); err != nil {
		switch {
		case errors.Is(err, domain.ErrServiceNotFound):
			return httpx.NotFound(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.NoContent(c)
}
