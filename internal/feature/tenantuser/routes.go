package tenantuser

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/tenantuser/handler"
)

type Routes struct {
	handler *handler.Handler
}

var _ router.RouteRegistrar = (*Routes)(nil)

func (r *Routes) RegisterRoutes(router fiber.Router) {
	tu := router.Group("tenant-users")

	tu.Get("/", r.handler.List)
	tu.Post("/", r.handler.Create)
	tu.Get("/:id", r.handler.Get)
	tu.Patch("/:id/status/:status", r.handler.UpdateStatus)
	tu.Delete("/:id", r.handler.Delete)
	tu.Delete("/:id/purge", r.handler.Purge)
}

func NewRoutes(handler *handler.Handler) *Routes {
	return &Routes{handler: handler}
}
