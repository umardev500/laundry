package servicecategory

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/servicecategory/handler"
)

// Routes defines all HTTP routes for the ServiceCategory feature.
type Routes struct {
	handler *handler.Handler
}

// Ensure Routes implements router.RouteRegistrar.
var _ router.RouteRegistrar = (*Routes)(nil)

// RegisterRoutes registers all endpoints for service categories.
func (r *Routes) RegisterRoutes(router fiber.Router) {
	categories := router.Group("service-categories")

	categories.Post("/", r.handler.Create)
	categories.Get("/", r.handler.List)
	categories.Get("/:id", r.handler.Get)
	categories.Put("/:id", r.handler.Update)
	categories.Delete("/:id", r.handler.Delete)
	categories.Delete("/:id/purge", r.handler.Purge)
}

// NewRoutes creates a new Routes instance.
func NewRoutes(h *handler.Handler) *Routes {
	return &Routes{
		handler: h,
	}
}
