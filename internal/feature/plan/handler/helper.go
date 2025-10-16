package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/plan/domain"
	"github.com/umardev500/laundry/pkg/httpx"
)

// handlePlanError centralizes common error handling for plan operations.
func handlePlanError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domain.ErrPlanDeleted),
		errors.Is(err, domain.ErrUnauthorizedPlanAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, domain.ErrPlanNotFound):
		return httpx.NotFound(c, err.Error())

	case errors.Is(err, domain.ErrPlanAlreadyExists):
		return httpx.Conflict(c, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
