package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/tenant/contract"
	"github.com/umardev500/laundry/internal/feature/tenant/domain"
	"github.com/umardev500/laundry/internal/feature/tenant/dto"
	"github.com/umardev500/laundry/internal/feature/tenant/mapper"
	"github.com/umardev500/laundry/internal/feature/tenant/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"
)

type Handler struct {
	service   contract.Service
	validator *validator.Validator
}

func NewHandler(service contract.Service, validator *validator.Validator) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

// 🏗️ Create a new Tenant
func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateTenantRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	tenantDomain, err := req.ToDomain()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	result, err := h.service.Create(ctx, tenantDomain)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrTenantAlreadyExists):
			return httpx.Conflict(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToTenantResponse(result))
}

// 🔍 Get a Tenant by ID
func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	result, err := h.service.GetByID(ctx, id)
	if err != nil {
		return handleTenantError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToTenantResponse(result))
}

// ⚙️ UpdateStatus changes tenant status
func (h *Handler) UpdateStatus(c *fiber.Ctx) error {
	var q query.UpdateStatusQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := q.Validate(); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	id, err := q.UUID()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	t := q.ToDomainTenantWithID(id)

	result, err := h.service.UpdateStatus(ctx, t)
	if err != nil {
		return handleTenantError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToTenantResponse(result))
}

// 🗑️ Soft Delete a Tenant
func (h *Handler) Delete(c *fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Delete(ctx, id)
	if err != nil {
		return handleTenantError(c, err)
	}

	return httpx.NoContent(c)
}

// 💣 Permanently delete (purge) a Tenant
func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Purge(ctx, id)
	if err != nil {
		return handleTenantError(c, err)
	}

	return httpx.NoContent(c)
}

// 📄 List tenants (with pagination, search, status, order)
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListTenantQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())
	result, err := h.service.List(ctx, &q)
	if err != nil {
		return handleTenantError(c, err)
	}

	dtoPage := mapper.ToTenantResponsePage(result)
	return httpx.JSONPaginated(c, fiber.StatusOK, dtoPage.Data, httpx.NewPagination(q.Page, q.Limit, result.Total))
}
