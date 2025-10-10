package handler

import (
	"errors"

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

// PermissionHandler handles HTTP requests for Permission resources.
type PermissionHandler struct {
	service   contract.PermissionService
	validator *validator.Validator
}

// NewPermissionHandler creates a new PermissionHandler instance.
func NewPermissionHandler(service contract.PermissionService, validator *validator.Validator) *PermissionHandler {
	return &PermissionHandler{
		service:   service,
		validator: validator,
	}
}

// 🔍 Get a Permission by ID
func (h *PermissionHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	result, err := h.service.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPermissionNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrPermissionDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToPermissionResponse(result))
}

// ⚙️ Update an existing Permission
func (h *PermissionHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdatePermissionRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	permDomain, err := req.ToDomain(id)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	result, err := h.service.Update(ctx, permDomain)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPermissionNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrPermissionDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToPermissionResponse(result))
}

// 🔄 Update a Permission's status (e.g., activate/suspend)
func (h *PermissionHandler) UpdateStatus(c *fiber.Ctx) error {
	var q query.UpdatePermissionStatusQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := q.Validate(); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	permDomain, err := q.ToDomain()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	result, err := h.service.UpdateStatus(ctx, permDomain.ID, permDomain)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPermissionNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrPermissionDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToPermissionResponse(result))
}

// 🗑️ Soft Delete a Permission
func (h *PermissionHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Delete(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPermissionDeleted):
			return httpx.Forbidden(c, err.Error())
		case errors.Is(err, domain.ErrPermissionNotFound):
			return httpx.NotFound(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.NoContent(c)
}

// 💣 Permanently delete (purge) a Permission
func (h *PermissionHandler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Purge(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPermissionNotFound):
			return httpx.NotFound(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.NoContent(c)
}

// 📄 List Permissions (pagination + search)
func (h *PermissionHandler) List(c *fiber.Ctx) error {
	var q query.ListPermissionQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())
	result, err := h.service.List(ctx, &q)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	dtoPage := mapper.ToPermissionResponsePage(result)
	return httpx.JSONPaginated(c, fiber.StatusOK, dtoPage.Data, httpx.NewPagination(q.Page, q.Limit, result.Total))
}
