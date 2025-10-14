package payment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/middleware"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/payment/handler"
)

type Routes struct {
	handler *handler.Handler
	config  *config.Config
}

var _ router.RouteRegistrar = (*Routes)(nil)

func (r *Routes) RegisterRoutes(router fiber.Router) {
	group := router.Group("payments")

	group.Use(middleware.CheckAuth(r.config))
	group.Get("/", r.handler.List)
	group.Get("/:id", r.handler.FindById)
}

func NewRoutes(h *handler.Handler, config *config.Config) *Routes {
	return &Routes{
		handler: h,
		config:  config,
	}
}
