package rbac

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/rbac/handler"
)

// Routes holds the handler for Role endpoints.
type Routes struct {
	handler *handler.Handler
}

// Ensure Routes implements the RouteRegistrar interface
var _ router.RouteRegistrar = (*Routes)(nil)

// RegisterRoutes implements router.RouteRegistrar.
func (r *Routes) RegisterRoutes(router fiber.Router) {
	role := router.Group("roles")

	role.Post("/", r.handler.Create)           // Create a new role
	role.Get("/", r.handler.List)              // List roles (with pagination, filters)
	role.Get("/:id", r.handler.Get)            // Get role by ID
	role.Put("/:id", r.handler.Update)         // Update role
	role.Delete("/:id", r.handler.Delete)      // Soft delete a role
	role.Delete("/:id/purge", r.handler.Purge) // Hard delete a role (permanent)
}

// NewRoutes creates a new Role routes instance.
func NewRoutes(handler *handler.Handler) *Routes {
	return &Routes{
		handler: handler,
	}
}
