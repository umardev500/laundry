package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/order/handler"
)

// Routes defines all HTTP routes for the Order feature.
type Routes struct {
	handler *handler.Handler
}

// Ensure Routes implements router.RouteRegistrar.
var _ router.RouteRegistrar = (*Routes)(nil)

// RegisterRoutes registers all endpoints for orders.
func (r *Routes) RegisterRoutes(router fiber.Router) {
	orders := router.Group("orders")

	// Since currently we only have List functionality
	orders.Get("/", r.handler.List)
	orders.Post("/guest", r.handler.GuestOrder)

	// If more handlers like Get, Create, Update, Delete are added later, register here
	// orders.Get("/:id", r.handler.Get)
	// orders.Put("/:id", r.handler.Update)
	// orders.Delete("/:id", r.handler.Delete)
}

// NewRoutes creates a new Routes instance.
func NewRoutes(h *handler.Handler) *Routes {
	return &Routes{
		handler: h,
	}
}
