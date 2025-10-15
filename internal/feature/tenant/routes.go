package tenant

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/middleware"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/tenant/handler"
)

// Routes holds the handler for tenant endpoints.
type Routes struct {
	handler *handler.Handler
	config  *config.Config
}

// Ensure Routes implements the RouteRegistrar interface
var _ router.RouteRegistrar = (*Routes)(nil)

// RegisterRoutes implements router.RouteRegistrar.
func (r *Routes) RegisterRoutes(router fiber.Router) {
	t := router.Group("tenants")

	t.Use(middleware.CheckAuth(r.config))
	t.Post("/", r.handler.Create)                          // Create a new tenant
	t.Get("/", r.handler.List)                             // List tenants (with pagination, filters)
	t.Get("/:id", r.handler.Get)                           // Get tenant by ID
	t.Delete("/:id", r.handler.Delete)                     // Soft delete
	t.Delete("/:id/purge", r.handler.Purge)                // Hard delete
	t.Patch("/:id/status/:status", r.handler.UpdateStatus) // Update tenant status
}

// NewRoutes creates a new tenant routes instance.
func NewRoutes(handler *handler.Handler, config *config.Config) *Routes {
	return &Routes{
		handler: handler,
		config:  config,
	}
}
