package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/auth/contract"
	"github.com/umardev500/laundry/internal/feature/auth/domain"
	"github.com/umardev500/laundry/internal/feature/auth/dto"
	"github.com/umardev500/laundry/internal/feature/auth/mapper"
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

// Login handles POST /auth/login
func (h *Handler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpx.BadRequest(c, err.Error())
	}

	ctx := appctx.New(c.UserContext())

	result, err := h.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidCredentials):
			return httpx.Unauthorized(c, err.Error())
		default:
			return httpx.InternalServerError(c, err.Error())
		}
	}

	// Map domain to DTO
	resp := mapper.FromDomain(result)

	return httpx.JSON(c, fiber.StatusOK, resp)
}
