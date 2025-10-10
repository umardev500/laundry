package machinetype

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/machinetype/handler"
)

type Routes struct {
	handler *handler.Handler
}

var _ router.RouteRegistrar = (*Routes)(nil)

func (r *Routes) RegisterRoutes(router fiber.Router) {
	m := router.Group("machine-types")

	m.Post("/", r.handler.Create)
	m.Get("/", r.handler.List)
	m.Get("/:id", r.handler.Get)
	m.Put("/:id", r.handler.Update)
	m.Delete("/:id", r.handler.Delete)
	m.Delete("/:id/purge", r.handler.Purge)
}

func NewRoutes(handler *handler.Handler) *Routes {
	return &Routes{handler: handler}
}
