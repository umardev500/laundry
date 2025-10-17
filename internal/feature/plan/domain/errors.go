package domain

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/pkg/httpx"
)

type PlanError struct {
	Err error
}

func (e PlanError) Error() string {
	return e.Err.Error()
}

// Optional: helper for errors.Is to work with the wrapper
func (e PlanError) Unwrap() error {
	return e.Err
}

var (
	ErrPlanAlreadyExists      = fmt.Errorf("plan already exists")
	ErrPlanNotFound           = fmt.Errorf("plan not found")
	ErrPlanNotActive          = fmt.Errorf("plan is not active")
	ErrPlanDeleted            = fmt.Errorf("plan has been deleted")
	ErrUnauthorizedPlanAccess = fmt.Errorf("unauthorized access to plan")
)

func NewPlanError(err error) PlanError {
	return PlanError{Err: err}
}

// HandlePlanError centralizes common error handling for plan operations.
func HandlePlanError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, ErrPlanNotActive):
		return httpx.BadRequest(c, err.Error())

	case errors.Is(err, ErrPlanDeleted),
		errors.Is(err, ErrUnauthorizedPlanAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, ErrPlanNotFound):
		return httpx.NotFound(c, err.Error())

	case errors.Is(err, ErrPlanAlreadyExists):
		return httpx.Conflict(c, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
