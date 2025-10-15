package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/servicecategory/domain"
	"github.com/umardev500/laundry/pkg/httpx"
)

// handleServiceCategoryError centralizes HTTP error mapping for the service category module.
func handleServiceCategoryError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, domain.ErrServiceCategoryDeleted),
		errors.Is(err, domain.ErrUnauthorizedAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, domain.ErrServiceCategoryNotFound):
		return httpx.NotFound(c, err.Error())

	case errors.Is(err, domain.ErrServiceCategoryAlreadyExists):
		return httpx.Conflict(c, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
