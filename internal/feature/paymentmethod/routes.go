package paymentmethod

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/handler"
)

type Routes struct {
	handler *handler.Handler
}

var _ router.RouteRegistrar = (*Routes)(nil)

func (r *Routes) RegisterRoutes(router fiber.Router) {
	group := router.Group("payment-methods")
	group.Post("/", r.handler.Create)
	group.Get("/", r.handler.List)
	group.Get("/:id", r.handler.Get)
	group.Put("/:id", r.handler.Update)
	group.Delete("/:id", r.handler.Delete)
	group.Delete("/:id/purge", r.handler.Purge)
}

func NewRoutes(h *handler.Handler) *Routes {
	return &Routes{handler: h}
}
