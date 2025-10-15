package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/serviceunit/domain"
	"github.com/umardev500/laundry/pkg/httpx"
)

// handleServiceUnitError centralizes common error handling for service unit operations
func handleServiceUnitError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domain.ErrServiceUnitDeleted),
		errors.Is(err, domain.ErrUnauthorizedAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, domain.ErrServiceUnitNotFound):
		return httpx.NotFound(c, err.Error())

	case errors.Is(err, domain.ErrServiceUnitAlreadyExists):
		return httpx.Conflict(c, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
