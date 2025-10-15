package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/rbac/contract"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/dto"
	"github.com/umardev500/laundry/internal/feature/rbac/mapper"
	"github.com/umardev500/laundry/internal/feature/rbac/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"
)

// Handler handles HTTP requests for Role resources.
type Handler struct {
	service   contract.Service
	validator *validator.Validator
}

// NewHandler creates a new Role handler instance.
func NewHandler(service contract.Service, validator *validator.Validator) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

// 🏗️ Create a new Role
func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	roleDomain, err := req.ToDomain(ctx.TenantID())
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	result, err := h.service.Create(ctx, roleDomain)
	if err != nil {
		return handleRoleError(c, err)
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToRoleResponse(result))
}

// 🔍 Get a Role by ID
func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	result, err := h.service.GetByID(ctx, id)
	if err != nil {
		return handleRoleError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToRoleResponse(result))
}

// ⚙️ Update an existing Role
func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdateRoleRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())

	roleDomain := &domain.Role{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	result, err := h.service.Update(ctx, roleDomain)
	if err != nil {
		return handleRoleError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToRoleResponse(result))
}

// 🗑️ Soft Delete a Role
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Delete(ctx, id)
	if err != nil {
		return handleRoleError(c, err)
	}

	return httpx.NoContent(c)
}

// 💣 Permanently delete (purge) a Role
func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Purge(ctx, id)
	if err != nil {
		return handleRoleError(c, err)
	}

	return httpx.NoContent(c)
}

// 📄 List Roles (with pagination, search, and ordering)
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListRoleQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())

	result, err := h.service.List(ctx, &q)
	if err != nil {
		return handleRoleError(c, err)
	}

	dtoPage := mapper.ToRoleResponsePage(result)
	return httpx.JSONPaginated(c, fiber.StatusOK, dtoPage.Data, httpx.NewPagination(q.Page, q.Limit, result.Total))
}
