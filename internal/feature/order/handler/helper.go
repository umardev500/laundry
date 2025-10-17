package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/order/dto"
	"github.com/umardev500/laundry/pkg/errorsx"
	"github.com/umardev500/laundry/pkg/httpx"
	"github.com/umardev500/laundry/pkg/types"

	paymentDomain "github.com/umardev500/laundry/internal/feature/payment/domain"
)

// handleOrderError centralizes HTTP error mapping for order module
func handleOrderError(c *fiber.Ctx, err error) error {
	var svcErr *domain.ServiceUnavailableError
	isServiceUnavailable := errors.As(err, &svcErr)

	switch {
	case errorsx.IsInvalidTransitionErr[types.OrderStatus](err):
		return httpx.JSONErrorWithData(
			c,
			fiber.StatusBadRequest,
			"invalid status transition",
			err,
			err,
		)

	case errors.Is(err, types.ErrStatusUnchanged):
		return httpx.JSONWithMessage[*dto.OrderResponse](c, fiber.StatusOK, nil, err.Error())

	case errors.Is(err, domain.ErrOrderDeleted),
		errors.Is(err, domain.ErrUnauthorizedOrderAccess):
		return httpx.Forbidden(c, err.Error())

	case errors.Is(err, domain.ErrOrderNotFound):
		return httpx.NotFound(c, err.Error())

	case isServiceUnavailable,
		errors.Is(err, domain.ErrGuestEmailOrPhoneRequired),
		errors.Is(err, domain.ErrOrderItemsRequired),
		errors.Is(err, paymentDomain.ErrInsufficientPayment):
		return httpx.BadRequest(c, err.Error())

	default:
		return httpx.InternalServerError(c, err.Error())
	}
}
