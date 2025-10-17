package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/region/contract"
	"github.com/umardev500/laundry/internal/feature/region/mapper"
	"github.com/umardev500/laundry/internal/feature/region/query"
	"github.com/umardev500/laundry/pkg/httpx"
)

type Handler struct {
	service contract.Service
}

func NewHandler(service contract.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ListProvince(c *fiber.Ctx) error {
	var q query.ListProvinceQuery
	if err := c.QueryParser(&q); err != nil {
		return err
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())
	data, err := h.service.FindProvinces(ctx, &q)
	if err != nil {
		return err
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		mapper.ToResponseProvinceList(data.Data),
		httpx.NewPagination(q.Page, q.Limit, data.Total),
	)
}

func (h *Handler) ListRegency(c *fiber.Ctx) error {
	id := c.Params("province_id")
	if id == "" {
		return httpx.BadRequest(c, fmt.Errorf("id is required").Error())
	}

	var q query.ListRegencyQuery
	if err := c.QueryParser(&q); err != nil {
		return err
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())
	data, err := h.service.FindRegencies(ctx, id, &q)
	if err != nil {
		return err
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		mapper.ToResponseRegencyList(data.Data),
		httpx.NewPagination(q.Page, q.Limit, data.Total),
	)
}

func (h *Handler) ListDistrict(c *fiber.Ctx) error {
	id := c.Params("regency_id")
	if id == "" {
		return httpx.BadRequest(c, fmt.Errorf("id is required").Error())
	}

	var q query.ListDistrictQuery
	if err := c.QueryParser(&q); err != nil {
		return err
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())
	data, err := h.service.FindDistricts(ctx, id, &q)
	if err != nil {
		return err
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		mapper.ToResponseDistrictList(data.Data),
		httpx.NewPagination(q.Page, q.Limit, data.Total),
	)
}

func (h *Handler) ListVillage(c *fiber.Ctx) error {
	id := c.Params("district_id")
	if id == "" {
		return httpx.BadRequest(c, fmt.Errorf("id is required").Error())
	}

	var q query.ListVillageQuery
	if err := c.QueryParser(&q); err != nil {
		return err
	}
	q.Normalize()

	ctx := appctx.New(c.UserContext())
	data, err := h.service.FindVillages(ctx, id, &q)
	if err != nil {
		return err
	}

	return httpx.JSONPaginated(
		c,
		fiber.StatusOK,
		mapper.ToResponseVillageList(data.Data),
		httpx.NewPagination(q.Page, q.Limit, data.Total),
	)
}
