package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/feature/user/handler"
)

type Routes struct {
	handler *handler.Handler
}

// RegisterRoutes implements router.RouteRegistrar.
func (r *Routes) RegisterRoutes(router fiber.Router) {
	user := router.Group("users")

	user.Post("/", r.handler.Create)
	user.Get("/", r.handler.List)
	user.Get("/:id", r.handler.GetUser)
	user.Delete("/:id", r.handler.Delete)
	user.Delete("/:id/purge", r.handler.Purge)
	user.Put("/:id", r.handler.Update)
	user.Patch("/:id/status/:status", r.handler.UpdateStatus)
}

func NewRoutes(handler *handler.Handler) router.RouteRegistrar {
	return &Routes{
		handler: handler,
	}
}
