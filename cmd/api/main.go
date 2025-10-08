package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/di"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

}

func main() {
	cfg := config.LoadConfig("./config/config.yml")
	client := entdb.NewEntClient(cfg)
	defer client.Client.Close()

	switch cfg.App.Env {
	case "development":
		log.Info().Msg("🧩 Running in development mode")
	case "staging":
		log.Info().Msg("🧪 Running in staging mode")
	case "production":
		log.Info().Msg("🚀 Running in production mode")
	default:
		log.Info().Msgf("⚙️ Running in %s mode", cfg.App.Env)
	}

	ctx := context.Background()
	if err := entdb.RunMigration(ctx, client.Client); err != nil {
		log.Fatal().Err(err).Msg("Failed to run migration")
	}

	router, err := di.Initialize(cfg)
	if err != nil {
		panic(err)
	}

	go func() {
		if err := router.Run(); err != nil {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := router.Shutdown(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Failed to shutdown server")
	}
}
