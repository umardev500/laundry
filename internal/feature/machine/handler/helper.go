package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/machine/domain"
	"github.com/umardev500/laundry/internal/feature/machine/dto"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/types"
)

// handleMachineError centralizes HTTP error mapping for machine module
func handleMachineError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domain.ErrMachineDeleted),
		errors.Is(err, domain.ErrUnauthorizedMachineAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, domain.ErrMachineNotFound):
		return httpx.NotFound(c, err.Error())

	case errors.Is(err, domain.ErrMachineAlreadyExists):
		return httpx.Conflict(c, err.Error())

	case errors.Is(err, types.ErrStatusUnchanged):
		return httpx.JSONWithMessage[*dto.MachineResponse](c, fiber.StatusOK, nil, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
