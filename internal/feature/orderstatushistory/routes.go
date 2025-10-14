package orderstatushistory

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/handler"
)

type Routes struct {
	handler *handler.Handler
}

var _ router.RouteRegistrar = (*Routes)(nil)

func (r *Routes) RegisterRoutes(router fiber.Router) {
	group := router.Group("order-status-history")
	group.Get("/", r.handler.List)
	group.Get("/:id", r.handler.GetByID)
}

func NewRoutes(h *handler.Handler) *Routes {
	return &Routes{handler: h}
}
