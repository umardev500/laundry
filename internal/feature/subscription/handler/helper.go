package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/subscription/domain"

	planDomain "github.com/umardev500/laundry/internal/feature/plan/domain"
)

// handleSubscriptionError centralizes error handling for subscription operations.
func handleSubscriptionError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case planDomain.PlanError:
		return planDomain.HandlePlanError(c, e)
	default:
		return domain.HandleSubscriptionError(c, e)
	}
}
