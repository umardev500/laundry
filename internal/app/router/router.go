package router

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

type RouteRegistrar interface {
	RegisterRoutes(router fiber.Router)
}

type Router struct {
	App    *fiber.App
	Client *entdb.Client
	Config *config.Config
}

func NewRouter(app *fiber.App, cfg *config.Config, client *entdb.Client, registrars []RouteRegistrar) *Router {
	api := app.Group("api")

	for _, r := range registrars {
		r.RegisterRoutes(api)
	}
	return &Router{
		App:    app,
		Client: client,
		Config: cfg,
	}
}

func (a *Router) Run() error {
	addr := ":" + a.Config.Server.Port
	log.Printf("🚀 Server running on %s", addr)
	return a.App.Listen(addr)
}

func (a *Router) Shutdown(ctx context.Context) error {
	log.Info().Msg("Shutting down server")
	if err := a.Client.Client.Close(); err != nil {
		return err
	}
	return a.App.Shutdown()
}
