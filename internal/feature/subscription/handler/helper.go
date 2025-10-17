package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/subscription/domain"
	"github.com/umardev500/laundry/pkg/errorsx"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/types"
)

// handleSubscriptionError centralizes error handling for subscription operations.
func handleSubscriptionError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domain.ErrSubscriptionAlreadyExists):
		return httpx.BadRequest(c, err.Error())

	case errorsx.IsInvalidTransitionErr[types.SubscriptionStatus](err):
		return httpx.JSONErrorWithData(
			c,
			fiber.StatusBadRequest,
			"invalid status transition",
			err,
			err,
		)

	case errors.Is(err, domain.ErrSubscriptionDeleted),
		errors.Is(err, domain.ErrUnauthorizedSubscriptionAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, domain.ErrSubscriptionNotFound):
		return httpx.NotFound(c, err.Error())

	case errorsx.IsInvalidTransitionErr[types.SubscriptionStatus](err):
		return httpx.BadRequest(c, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
