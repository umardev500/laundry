package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/service/domain"
	"github.com/umardev500/laundry/pkg/httpx"
)

// handleServiceError centralizes HTTP error mapping for service module
func handleServiceError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domain.ErrServiceDeleted),
		errors.Is(err, domain.ErrUnauthorizedServiceAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, domain.ErrServiceNotFound):
		return httpx.NotFound(c, err.Error())

	case errors.Is(err, domain.ErrServiceAlreadyExists):
		return httpx.Conflict(c, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
