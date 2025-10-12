package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/middleware"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/auth/handler"
)

type Routes struct {
	handler *handler.Handler
	config  *config.Config
}

// Ensure Routes implements the RouteRegistrar interface
var _ router.RouteRegistrar = (*Routes)(nil)

// RegisterRoutes implements router.RouteRegistrar.
func (r *Routes) RegisterRoutes(router fiber.Router) {
	auth := router.Group("auth")

	// Login endpoint
	auth.Post("/login", r.handler.Login)

	// Add more auth routes here later (refresh token, logout, etc.)
	auth.Use(middleware.CheckAuth(r.config))
}

// NewRoutes returns a new Routes instance
func NewRoutes(handler *handler.Handler, config *config.Config) *Routes {
	return &Routes{
		handler: handler,
		config:  config,
	}
}
