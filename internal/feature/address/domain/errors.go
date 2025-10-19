package domain

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/pkg/httpx"
)

// --- Address-specific domain errors ---
var (
	ErrAddressNotFound        = fmt.Errorf("address not found")
	ErrAddressDeleted         = fmt.Errorf("address has been deleted")
	ErrAddressNotDeleted      = fmt.Errorf("address is not deleted")
	ErrAddressAlreadyExists   = fmt.Errorf("address already exists")
	ErrUnauthorizedAddress    = fmt.Errorf("unauthorized access to address")
	ErrInvalidAddressLocation = fmt.Errorf("invalid address location identifiers")
)

// AddressError wraps address-related errors for consistent handling.
type AddressError struct {
	Err error
}

func (e AddressError) Error() string {
	return e.Err.Error()
}

func (e AddressError) Unwrap() error {
	return e.Err
}

// NewAddressError wraps a given error into an AddressError.
func NewAddressError(err error) AddressError {
	return AddressError{Err: err}
}

// HandleAddressError centralizes HTTP error responses for address operations.
func HandleAddressError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, ErrAddressAlreadyExists):
		return httpx.BadRequest(c, err.Error())

	case errors.Is(err, ErrInvalidAddressLocation):
		return httpx.BadRequest(c, err.Error())

	case errors.Is(err, ErrAddressDeleted),
		errors.Is(err, ErrUnauthorizedAddress):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, ErrAddressNotFound):
		return httpx.NotFound(c, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
