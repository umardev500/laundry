package main

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

}

func main() {
	cfg := config.LoadConfig("./config/config.yml")
	client := entdb.NewEntClient(cfg)
	defer client.Client.Close()

	ctx := context.Background()
	if err := entdb.RunMigration(ctx, client.Client); err != nil {
		log.Fatal().Err(err).Msg("Failed to run migration")
	}
}
