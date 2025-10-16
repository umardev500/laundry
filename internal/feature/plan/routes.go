package plan

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/middleware"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/plan/handler"
)

type Routes struct {
	handler *handler.Handler
	config  *config.Config
}

var _ router.RouteRegistrar = (*Routes)(nil)

func (r *Routes) RegisterRoutes(router fiber.Router) {
	group := router.Group("plans")

	group.Use(middleware.CheckAuth(r.config))
	group.Post("/", r.handler.Create)
	group.Get("/", r.handler.List)
	group.Get("/:id", r.handler.Get)
	group.Put("/:id", r.handler.Update)
	group.Delete("/:id", r.handler.Delete)
	group.Delete("/:id/purge", r.handler.Purge)

	// Status endpoints
	group.Patch("/:id/activate", r.handler.Activate)
	group.Patch("/:id/deactivate", r.handler.Deactivate)
	group.Patch("/:id/restore", r.handler.Restore)
}

func NewRoutes(h *handler.Handler, config *config.Config) *Routes {
	return &Routes{
		handler: h,
		config:  config,
	}
}
