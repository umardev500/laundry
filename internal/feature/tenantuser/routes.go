package tenantuser

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/middleware"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/tenantuser/handler"
)

type Routes struct {
	handler *handler.Handler
	config  *config.Config
}

var _ router.RouteRegistrar = (*Routes)(nil)

func (r *Routes) RegisterRoutes(router fiber.Router) {
	tu := router.Group("tenant-users")

	tu.Use(middleware.CheckAuth(r.config))
	tu.Get("/", r.handler.List)
	tu.Post("/", r.handler.Create)
	tu.Get("/:id", r.handler.Get)
	tu.Patch("/:id/status/:status", r.handler.UpdateStatus)
	tu.Delete("/:id", r.handler.Delete)
	tu.Delete("/:id/purge", r.handler.Purge)
}

func NewRoutes(handler *handler.Handler, config *config.Config) *Routes {
	return &Routes{
		handler: handler,
		config:  config,
	}
}
