package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/machine/contract"
	"github.com/umardev500/laundry/internal/feature/machine/dto"
	"github.com/umardev500/laundry/internal/feature/machine/mapper"
	"github.com/umardev500/laundry/internal/feature/machine/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"
)

type Handler struct {
	service   contract.Service
	validator *validator.Validator
}

func NewHandler(service contract.Service, validator *validator.Validator) *Handler {
	return &Handler{service: service, validator: validator}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateMachineRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	mDomain, err := req.ToDomain(ctx)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	res, err := h.service.Create(ctx, mDomain)
	if err != nil {
		return handleMachineError(c, err)
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToMachineResponse(res))
}

func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	res, err := h.service.GetByID(ctx, id)
	if err != nil {
		return handleMachineError(c, err)
	}
	return httpx.JSON(c, fiber.StatusOK, mapper.ToMachineResponse(res))
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdateMachineRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	machineDomain, err := req.ToDomain(id)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	res, err := h.service.Update(ctx, machineDomain)
	if err != nil {
		return handleMachineError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToMachineResponse(res))
}

func (h *Handler) UpdateStatus(c *fiber.Ctx) error {
	var q query.UpdateStatusQuery
	if err := c.ParamsParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	m, err := q.ToDomain(id)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	res, err := h.service.UpdateStatus(ctx, m)
	if err != nil {
		return handleMachineError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToMachineResponse(res))
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Delete(ctx, id)
	if err != nil {
		return handleMachineError(c, err)
	}
	return httpx.NoContent(c)
}

func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())
	err = h.service.Purge(ctx, id)
	if err != nil {
		return handleMachineError(c, err)
	}
	return httpx.NoContent(c)
}

func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListMachineQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())
	res, err := h.service.List(ctx, &q)
	if err != nil {
		return handleMachineError(c, err)
	}

	dtoPage := mapper.ToMachineResponsePage(res)
	return httpx.JSONPaginated(c, fiber.StatusOK, dtoPage.Data, httpx.NewPagination(q.Page, q.Limit, res.Total))
}
