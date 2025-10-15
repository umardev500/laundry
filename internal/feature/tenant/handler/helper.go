package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/tenant/domain"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/types"
)

// handleTenantError centralizes HTTP error mapping for tenant module
func handleTenantError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domain.ErrTenantDeleted),
		errors.Is(err, domain.ErrUnauthorizedTenant):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, domain.ErrTenantNotFound):
		return httpx.NotFound(c, err.Error())

	case errors.Is(err, domain.ErrTenantAlreadyExists):
		return httpx.Conflict(c, err.Error())

	case errors.Is(err, types.ErrStatusUnchanged):
		return httpx.JSONWithMessage[any](c, fiber.StatusOK, nil, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
