package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/tenantuser/contract"
	"github.com/umardev500/laundry/internal/feature/tenantuser/dto"
	"github.com/umardev500/laundry/internal/feature/tenantuser/mapper"
	"github.com/umardev500/laundry/internal/feature/tenantuser/query"
	"github.com/umardev500/laundry/pkg/httpx"
	pkgQuery "github.com/umardev500/laundry/pkg/query"
	"github.com/umardev500/laundry/pkg/validator"
)

// Handler handles HTTP requests for tenant user management.
type Handler struct {
	service   contract.Service
	validator *validator.Validator
}

// NewTenantUserHandler creates a new TenantUserHandler instance.
func NewTenantUserHandler(service contract.Service, validator *validator.Validator) *Handler {
	return &Handler{service: service, validator: validator}
}

// 🔹 Create a TenantUser (assign a user to a tenant)
func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateTenantUserRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	tu := req.ToDomain(ctx)
	result, err := h.service.Create(ctx, tu)
	if err != nil {
		return handleTenantUserError(c, err)
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToTenantUserResponse(result))
}

// 📄 List all tenant users (paginated)
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListTenantUserQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())

	result, err := h.service.List(ctx, &q)
	if err != nil {
		return handleTenantUserError(c, err)
	}

	dtoPage := mapper.ToTenantUserResponsePage(result)
	return httpx.JSONPaginated(c, fiber.StatusOK, dtoPage.Data, httpx.NewPagination(q.Page, q.Limit, result.Total))
}

// 🔍 Get a tenant user by ID
func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	result, err := h.service.GetByID(ctx, id)
	if err != nil {
		return handleTenantUserError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToTenantUserResponse(result))
}

// ⚙️ Update tenant user status (activate/suspend)
func (h *Handler) UpdateStatus(c *fiber.Ctx) error {
	var q query.UpdateStatusQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := q.Validate(); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())

	tuDomain, err := q.ToDomain()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	result, err := h.service.UpdateStatus(ctx, tuDomain)
	if err != nil {
		return handleTenantUserError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToTenantUserResponse(result))
}

// 🗑️ Soft delete a tenant user
func (h *Handler) Delete(c *fiber.Ctx) error {
	var q pkgQuery.GetByIDQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	id, err := q.UUID()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	if err := h.service.Delete(ctx, id); err != nil {
		return handleTenantUserError(c, err)
	}

	return httpx.NoContent(c)
}

// 💣 Purge (hard delete)
func (h *Handler) Purge(c *fiber.Ctx) error {
	var q pkgQuery.GetByIDQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	id, err := q.UUID()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	if err := h.service.Purge(ctx, id); err != nil {
		return handleTenantUserError(c, err)
	}

	return httpx.NoContent(c)
}
