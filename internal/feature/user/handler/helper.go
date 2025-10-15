package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/user/domain"
	"github.com/umardev500/laundry/pkg/httpx"

	errorsPkg "github.com/umardev500/laundry/pkg/errors"
	"github.com/umardev500/laundry/pkg/types"
)

// HandleUserError maps domain user errors to proper HTTP responses.
func handleUserError(c *fiber.Ctx, err error) error {
	switch {
	case errorsPkg.IsInvalidTransitionErr[types.UserStatus](err):
		return httpx.JSONErrorWithData(
			c,
			fiber.StatusBadRequest,
			"invalid status transition",
			err,
			err,
		)
	case errors.Is(err, domain.ErrUserDeleted),
		errors.Is(err, domain.ErrUserSuspended),
		errors.Is(err, domain.ErrUnauthorizedUserAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, domain.ErrUserNotFound):
		return httpx.NotFound(c, err.Error())

	case errors.Is(err, domain.ErrUserAlreadyExists):
		return httpx.Conflict(c, err.Error())

	case errors.Is(err, types.ErrStatusUnchanged):
		return httpx.JSONWithMessage[any](c, fiber.StatusOK, nil, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
