package machine

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/middleware"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/machine/handler"
)

type Routes struct {
	handler *handler.Handler
	config  *config.Config
}

var _ router.RouteRegistrar = (*Routes)(nil)

func (r *Routes) RegisterRoutes(router fiber.Router) {
	m := router.Group("machines")

	m.Use(middleware.CheckAuth(r.config))
	m.Post("/", r.handler.Create)
	m.Get("/", r.handler.List)
	m.Get("/:id", r.handler.Get)
	m.Delete("/:id", r.handler.Delete)
	m.Delete("/:id/purge", r.handler.Purge)
	m.Put("/:id", r.handler.Update)
	m.Patch("/:id/status/:status", r.handler.UpdateStatus)
}

func NewRoutes(handler *handler.Handler, config *config.Config) *Routes {
	return &Routes{
		handler: handler,
		config:  config,
	}
}
