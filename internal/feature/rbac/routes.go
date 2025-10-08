package rbac

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/rbac/handler"
)

// Routes holds the handler for Role endpoints.
type Routes struct {
	handler           *handler.Handler
	featureHandler    *handler.FeatureHandler
	permissionHandler *handler.PermissionHandler
}

// Ensure Routes implements the RouteRegistrar interface
var _ router.RouteRegistrar = (*Routes)(nil)

// RegisterRoutes implements router.RouteRegistrar.
func (r *Routes) RegisterRoutes(router fiber.Router) {
	// Base route
	base := router.Group("rbac")

	// --- Role Routes ---
	role := base.Group("roles")

	role.Post("/", r.handler.Create)           // Create a new role
	role.Get("/", r.handler.List)              // List roles (with pagination, filters)
	role.Get("/:id", r.handler.Get)            // Get role by ID
	role.Put("/:id", r.handler.Update)         // Update role
	role.Delete("/:id", r.handler.Delete)      // Soft delete a role
	role.Delete("/:id/purge", r.handler.Purge) // Hard delete a role (permanent)

	// --- Feature Routes ---
	feature := base.Group("features")
	feature.Get("/", r.featureHandler.List)
	feature.Get("/:id", r.featureHandler.Get)
	feature.Put("/:id", r.featureHandler.Update)
	feature.Patch("/:id/:status", r.featureHandler.UpdateStatus)

	// --- Permission Routes ---
	perm := base.Group("permissions")
	perm.Get("/", r.permissionHandler.List)
	perm.Get("/:id", r.permissionHandler.Get)
	perm.Put("/:id", r.permissionHandler.Update)
	perm.Patch("/:id/:status", r.permissionHandler.UpdateStatus) // e.g. /permissions/:id/active or /permissions/:id/suspended
	perm.Delete("/:id", r.permissionHandler.Delete)
	perm.Delete("/:id/purge", r.permissionHandler.Purge)

}

// NewRoutes creates a new Role routes instance.
func NewRoutes(handler *handler.Handler, featureHandler *handler.FeatureHandler, permissionHandler *handler.PermissionHandler) *Routes {
	return &Routes{
		handler:           handler,
		featureHandler:    featureHandler,
		permissionHandler: permissionHandler,
	}
}
