package platformuser

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/platformuser/handler"
)

type Routes struct {
	handler *handler.Handler
}

// Ensure Routes implements the RouteRegistrar interface
var _ router.RouteRegistrar = (*Routes)(nil)

// RegisterRoutes implements router.RouteRegistrar.
func (r *Routes) RegisterRoutes(router fiber.Router) {
	pu := router.Group("platform-users")

	pu.Post("/", r.handler.Create)
	pu.Get("/", r.handler.List)
	pu.Get("/:id", r.handler.Get)
	pu.Delete("/:id", r.handler.Delete)
	pu.Delete("/:id/purge", r.handler.Purge)
	pu.Patch("/:id/status/:status", r.handler.UpdateStatus)
}

// NewRoutes returns a new Routes instance
func NewRoutes(handler *handler.Handler) *Routes {
	return &Routes{
		handler: handler,
	}
}
