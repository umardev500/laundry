package domain

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/pkg/errorsx"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/types"
)

var (
	// Subscription errors
	ErrSubscriptionNotFound           = fmt.Errorf("subscription not found")
	ErrSubscriptionDeleted            = fmt.Errorf("subscription has been deleted")
	ErrUnauthorizedSubscriptionAccess = fmt.Errorf("unauthorized access to subscription")
	ErrSubscriptionAlreadyExists      = fmt.Errorf("subscription already exists")
	ErrSubscriptionNotActive          = fmt.Errorf("subscription is not active")
	ErrSubscriptionNotDeleted         = fmt.Errorf("subscription is not deleted")
)

type SubscriptionError struct {
	Err error
}

func (e SubscriptionError) Error() string {
	return e.Err.Error()
}

// Optional: helper for errors.Is to work with the wrapper
func (e SubscriptionError) Unwrap() error {
	return e.Err
}

func NewSubscriptionError(err error) SubscriptionError {
	return SubscriptionError{Err: err}
}

// HandleSubscriptionError centralizes error handling for subscription operations.
func HandleSubscriptionError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, ErrSubscriptionAlreadyExists):
		return httpx.BadRequest(c, err.Error())

	case errorsx.IsInvalidTransitionErr[types.SubscriptionStatus](err):
		return httpx.JSONErrorWithData(
			c,
			fiber.StatusBadRequest,
			"invalid status transition",
			err,
			err,
		)

	case errors.Is(err, ErrSubscriptionDeleted),
		errors.Is(err, ErrUnauthorizedSubscriptionAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, ErrSubscriptionNotFound):
		return httpx.NotFound(c, err.Error())

	case errorsx.IsInvalidTransitionErr[types.SubscriptionStatus](err):
		return httpx.BadRequest(c, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
