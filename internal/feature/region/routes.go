package region

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/region/handler"
)

type Routes struct {
	handler *handler.Handler
	config  *config.Config
}

var _ router.RouteRegistrar = (*Routes)(nil)

func (r *Routes) RegisterRoutes(router fiber.Router) {
	group := router.Group("region")

	group.Get("/provinces", r.handler.ListProvince)
	group.Get("/regencies/:province_id", r.handler.ListRegency)
	group.Get("/districts/:regency_id", r.handler.ListDistrict)
	group.Get("/villages/:district_id", r.handler.ListVillage)
}

func NewRoutes(h *handler.Handler, config *config.Config) *Routes {
	return &Routes{
		handler: h,
		config:  config,
	}
}
