package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/servicecategory/contract"
	"github.com/umardev500/laundry/internal/feature/servicecategory/dto"
	"github.com/umardev500/laundry/internal/feature/servicecategory/mapper"
	"github.com/umardev500/laundry/internal/feature/servicecategory/query"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/validator"
)

type Handler struct {
	service   contract.Service
	validator *validator.Validator
}

func NewHandler(s contract.Service, v *validator.Validator) *Handler {
	return &Handler{service: s, validator: v}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateServiceCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	category, err := req.ToDomain(ctx)
	if err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	res, err := h.service.Create(ctx, category)
	if err != nil {
		return handleServiceCategoryError(c, err)
	}

	return httpx.JSON(c, fiber.StatusCreated, mapper.ToResponse(res))
}

func (h *Handler) List(c *fiber.Ctx) error {
	var q query.ListServiceCategoryQuery
	if err := c.QueryParser(&q); err != nil {
		return httpx.BadRequest(c, err.Error())
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())

	data, err := h.service.List(ctx, &q)
	if err != nil {
		return handleServiceCategoryError(c, err)
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		mapper.ToResponseList(data.Data),
		httpx.NewPagination(q.Page, q.Limit, data.Total),
	)
}

func (h *Handler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	category, err := h.service.GetByID(ctx, id)
	if err != nil {
		return handleServiceCategoryError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(category))
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	var req dto.UpdateServiceCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())
	category := req.ToDomain(id)

	res, err := h.service.Update(ctx, category)
	if err != nil {
		return handleServiceCategoryError(c, err)
	}

	return httpx.JSON(c, fiber.StatusOK, mapper.ToResponse(res))
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	if err := h.service.Delete(ctx, id); err != nil {
		return handleServiceCategoryError(c, err)
	}

	return httpx.NoContent(c)
}

func (h *Handler) Purge(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return httpx.BadRequest(c, "invalid id")
	}

	ctx := appctx.New(c.UserContext())

	if err := h.service.Purge(ctx, id); err != nil {
		return handleServiceCategoryError(c, err)
	}

	return httpx.NoContent(c)
}
