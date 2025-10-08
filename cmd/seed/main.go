package main

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/di"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig("./config/config.yml")

	seeders, err := di.InitialzeSeeder(cfg)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize seeders: %v", err)
	}

	// Run all seeders
	if err := seeder.RunAll(ctx, seeders); err != nil {
		log.Fatal().Err(err).Msg("Failed to run seeders")
	}
}
