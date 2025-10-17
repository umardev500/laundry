package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/tenantuser/domain"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/types"

	errorsPkg "github.com/umardev500/laundry/pkg/errorsx"
)

// handleTenantUserError maps domain-level tenant user errors to proper HTTP responses.
func handleTenantUserError(c *fiber.Ctx, err error) error {
	var transitionErr *errorsPkg.ErrInvalidStatusTransition[types.OrderStatus]
	isTransitionError := errors.As(err, &transitionErr)

	switch {
	case isTransitionError:
		return httpx.JSONErrorWithData(
			c,
			fiber.StatusBadRequest,
			"invalid status transition",
			transitionErr,
			err,
		)

	case errors.Is(err, domain.ErrTenantUserDeleted),
		errors.Is(err, domain.ErrUnauthorizedUserAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, domain.ErrTenantUserNotFound),
		errors.Is(err, domain.ErrTenantOrUserNotFound):
		return httpx.NotFound(c, err.Error())

	case errors.Is(err, domain.ErrTenantUserAlreadyExists):
		return httpx.Conflict(c, err.Error())

	case errors.Is(err, domain.ErrTenantIDMismatch):
		return httpx.BadRequest(c, err.Error())

	case errors.Is(err, domain.ErrInvalidStatusTransition):
		return httpx.UnprocessableEntity(c, err.Error())

	case errors.Is(err, types.ErrStatusUnchanged):
		return httpx.JSONWithMessage[any](c, fiber.StatusOK, nil, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
