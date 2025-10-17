package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/feature/plan/domain"
)

// handlePlanError centralizes common error handling for plan operations.
func handlePlanError(c *fiber.Ctx, err error) error {
	return domain.HandlePlanError(c, err)
}
