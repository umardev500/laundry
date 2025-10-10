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
	"github.com/umardev500/laundry/pkg/types"
	"github.com/umardev500/laundry/pkg/validator"
)

// FeatureHandler handles HTTP requests for Feature resources.
type FeatureHandler struct {
	service   contract.FeatureService
	validator *validator.Validator
}

// NewFeatureHandler creates a new FeatureHandler.
func NewFeatureHandler(service contract.FeatureService, validator *validator.Validator) *FeatureHandler {
	return &FeatureHandler{
		service:   service,
		validator: validator,
	}
}

// 🔍 Get a Feature by ID
func (h *FeatureHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	result, err := h.service.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrFeatureNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrFeatureDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToFeatureResponse(result))
}

// ⚙️ Update a Feature
func (h *FeatureHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdateFeatureRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	featureDomain, err := req.ToDomain(id)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	result, err := h.service.Update(ctx, featureDomain)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrFeatureNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrFeatureDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToFeatureResponse(result))
}

func (h *FeatureHandler) UpdateStatus(c *fiber.Ctx) error {
	var q query.UpdateFeatureStatusQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := q.Validate(); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	featureDomain, err := q.ToDomain()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	result, err := h.service.UpdateStatus(ctx, featureDomain)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrFeatureNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, types.ErrStatusUnchanged),
			errors.Is(err, types.ErrInvalidStatusTransition),
			errors.Is(err, types.ErrInvalidStatus):
			return httpx.BadRequest(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToFeatureResponse(result))
}

// 📄 List Features (pagination + search)
func (h *FeatureHandler) List(c *fiber.Ctx) error {
	var q query.ListFeatureQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())
	result, err := h.service.List(ctx, &q)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	dtoPage := mapper.ToFeatureResponsePage(result)
	return httpx.JSONPaginated(c, fiber.StatusOK, dtoPage.Data, httpx.NewPagination(q.Page, q.Limit, result.Total))
}
