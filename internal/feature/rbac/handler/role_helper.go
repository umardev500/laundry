package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/pkg/httpx"
)

// handleRoleError maps role domain errors to appropriate HTTP responses.
func handleRoleError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domain.ErrRoleAlreadyExists):
		return httpx.Conflict(c, err.Error()) // 409 Conflict

	case errors.Is(err, domain.ErrRoleNotFound):
		return httpx.NotFound(c, err.Error()) // 404 Not Found

	case errors.Is(err, domain.ErrEmptyRoleName),
		errors.Is(err, domain.ErrMissingTenantID):
		return httpx.BadRequest(c, err.Error()) // 400 Bad Request

	case errors.Is(err, domain.ErrUnauthorizedRoleAccess),
		errors.Is(err, domain.ErrRoleDeleted):
		return httpx.Forbidden(c, err.Error()) // 403 Forbidden

	default:
		return httpx.InternalServerError(c, err.Error()) // 500 fallback
	}
}
