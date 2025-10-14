package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/user/contract"
	"github.com/umardev500/laundry/internal/feature/user/domain"
	"github.com/umardev500/laundry/internal/feature/user/dto"
	"github.com/umardev500/laundry/internal/feature/user/mapper"
	"github.com/umardev500/laundry/internal/feature/user/query"
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

func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	user, err := req.ToDomainUser()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	result, err := h.service.Create(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserAlreadyExists):
			return httpx.Conflict(c, err.Error())

		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, int(fiber.StatusCreated), mapper.ToUserResponse(result))
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserDeleted) {
			return httpx.Forbidden(c, err.Error())
		}

		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.NoContent(c)
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	var q query.GetuserQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	uid, err := q.UUID()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	result, err := h.service.GetByID(ctx, uid)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.JSON(c, int(fiber.StatusOK), mapper.ToUserResponse(result))
}

func (h *Handler) Update(c *fiber.Ctx) error {
	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	user, err := req.ToDomainUserWithID(id)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	result, err := h.service.Update(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserNotFound):
			return httpx.NotFound(c, err.Error())

		case errors.Is(err, domain.ErrUserDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, int(fiber.StatusCreated), mapper.ToUserResponse(result))
}

func (h *Handler) UpdateStatus(c *fiber.Ctx) error {
	var q query.UpdateStatusQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := q.Validate(); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	uid, err := q.UUID()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	user := q.ToDomainUserWithID(uid)

	result, err := h.service.UpdateStatus(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserNotFound):
			return httpx.NotFound(c, err.Error())

		case errors.Is(err, domain.ErrUserDeleted):
			return httpx.Forbidden(c, err.Error())

		case errors.Is(err, domain.ErrStatusUnchanged):
			// Idempotent: nothing changed, but still OK
			return httpx.JSONWithMessage[*dto.UserResponse](c, fiber.StatusOK, nil, err.Error())

		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToUserResponse(result))
}

func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListUserQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize() // set defaults if not provided

	ctx := appctx.New(c.UserContext())
	result, err := h.service.List(ctx, &q)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}

	dtoPage := mapper.ToUserResponsePage(result)

	return httpx.JSONPaginated(c, int(fiber.StatusOK), dtoPage.Data, httpx.NewPagination(q.Page, q.Limit, result.Total))
}

func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Purge(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return httpx.NotFound(c, err.Error())
		}

		return httpx.InternalServerError(c, err.Error())
	}

	return httpx.NoContent(c)
}
