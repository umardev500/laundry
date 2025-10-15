package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/platformuser/domain"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/types"

	errorsPkg "github.com/umardev500/laundry/pkg/errors"
)

// handlePlatformUserError maps domain-level errors to HTTP responses.
func handlePlatformUserError(c *fiber.Ctx, err error) error {
	switch {
	case errorsPkg.IsInvalidTransitionErr[types.PlatformUserStatus](err):
		return httpx.JSONErrorWithData(
			c,
			fiber.StatusBadRequest,
			"invalid status transition",
			err,
			err,
		)

	case errors.Is(err, domain.ErrPlatformUserAlreadyExists):
		return httpx.Conflict(c, err.Error()) // 409

	case errors.Is(err, domain.ErrPlatformUserNotFound):
		return httpx.NotFound(c, err.Error()) // 404

	case errors.Is(err, domain.ErrPlatformUserDeleted):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, types.ErrStatusUnchanged):
		return httpx.JSONWithMessage[any](c, fiber.StatusOK, nil, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error()) // 500 fallback
	}
}
