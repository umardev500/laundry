package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/platformuser/contract"
	"github.com/umardev500/laundry/internal/feature/platformuser/domain"
	"github.com/umardev500/laundry/internal/feature/platformuser/dto"
	"github.com/umardev500/laundry/internal/feature/platformuser/mapper"
	"github.com/umardev500/laundry/internal/feature/platformuser/query"
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

// Create a new PlatformUser
func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreatePlatformUserRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	pu, err := req.ToDomain()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	result, err := h.service.Create(ctx, pu)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPlatformUserAlreadyExists):
			return httpx.Conflict(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToPlatformUserResponse(result))
}

// GetByID retrieves a PlatformUser by its ID
func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	result, err := h.service.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPlatformUserNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrPlatformUserDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToPlatformUserResponse(result))
}

// UpdateStatus updates the status of a PlatformUser
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
	pu := q.ToDomainPlatformUserWithID(id)

	result, err := h.service.UpdateStatus(ctx, pu)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPlatformUserNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrPlatformUserDeleted):
			return httpx.Forbidden(c, err.Error())
		case errors.Is(err, domain.ErrStatusUnchanged):
			// Idempotent: nothing changed, still OK
			return httpx.JSONWithMessage[*dto.PlatformUserResponse](c, fiber.StatusOK, nil, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToPlatformUserResponse(result))
}

// Delete performs a soft delete
func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Delete(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPlatformUserDeleted):
			return httpx.Forbidden(c, err.Error())
		case errors.Is(err, domain.ErrPlatformUserNotFound):
			return httpx.NotFound(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.NoContent(c)
}

// Purge permanently deletes a PlatformUser
func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Purge(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrPlatformUserNotFound):
			return httpx.NotFound(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.NoContent(c)
}

// List returns paginated PlatformUsers
func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListPlatformUserQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())
	result, err := h.service.List(ctx, &q)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	dtoPage := mapper.ToPlatformUserResponsePage(result)
	return httpx.JSONPaginated(c, fiber.StatusOK, dtoPage.Data, httpx.NewPagination(q.Page, q.Limit, result.Total))
}
