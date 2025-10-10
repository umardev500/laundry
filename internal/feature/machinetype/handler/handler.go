package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/machinetype/contract"
	"github.com/umardev500/laundry/internal/feature/machinetype/domain"
	"github.com/umardev500/laundry/internal/feature/machinetype/dto"
	"github.com/umardev500/laundry/internal/feature/machinetype/mapper"
	"github.com/umardev500/laundry/internal/feature/machinetype/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"

	pkgQuery "github.com/umardev500/laundry/pkg/query"
)

type Handler struct {
	service   contract.Service
	validator *validator.Validator
}

func NewHandler(service contract.Service, validator *validator.Validator) *Handler {
	return &Handler{service: service, validator: validator}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateMachineTypeRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	ctx := appctx.New(c.UserContext())
	domainObj, err := req.ToDomain(ctx)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	res, err := h.service.Create(ctx, domainObj)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrMachineTypeAlreadyExists):
			return httpx.Conflict(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}
	return httpx.JSON(c, fiber.StatusCreated, mapper.ToResponse(res))
}

func (h *Handler) Get(c *fiber.Ctx) error {
	var q pkgQuery.GetByIDQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	id, err := q.UUID()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	ctx := appctx.New(c.UserContext())
	res, err := h.service.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrMachineTypeNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrMachineTypeDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}
	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

func (h *Handler) Update(c *fiber.Ctx) error {
	var q pkgQuery.GetByIDQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	id, err := q.UUID()
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	var req dto.UpdateMachineTypeRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	ctx := appctx.New(c.UserContext())
	d, err := req.ToDomain(id)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	res, err := h.service.Update(ctx, d)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrMachineTypeNotFound):
			return httpx.NotFound(c, err.Error())
		case errors.Is(err, domain.ErrMachineTypeDeleted):
			return httpx.Forbidden(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}
	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

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
	err = h.service.Delete(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrMachineTypeDeleted):
			return httpx.Forbidden(c, err.Error())
		case errors.Is(err, domain.ErrMachineTypeNotFound):
			return httpx.NotFound(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}
	return httpx.NoContent(c)
}

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
	err = h.service.Purge(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrMachineTypeNotFound):
			return httpx.NotFound(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}
	return httpx.NoContent(c)
}

func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListMachineTypeQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()
	ctx := appctx.New(c.UserContext())
	res, err := h.service.List(ctx, &q)
	if err != nil {
		return httpx.InternalServerError(c, err.Error())
	}
	dtoPage := mapper.ToResponsePage(res)
	return httpx.JSONPaginated(c, fiber.StatusOK, dtoPage.Data, httpx.NewPagination(q.Page, q.Limit, res.Total))
}
