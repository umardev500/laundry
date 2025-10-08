package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/auth/handler"
)

type Routes struct {
	handler *handler.Handler
}

// Ensure Routes implements the RouteRegistrar interface
var _ router.RouteRegistrar = (*Routes)(nil)

// RegisterRoutes implements router.RouteRegistrar.
func (r *Routes) RegisterRoutes(router fiber.Router) {
	auth := router.Group("auth")

	// Login endpoint
	auth.Post("/login", r.handler.Login)

	// Add more auth routes here later (refresh token, logout, etc.)
}

// NewRoutes returns a new Routes instance
func NewRoutes(handler *handler.Handler) *Routes {
	return &Routes{
		handler: handler,
	}
}
