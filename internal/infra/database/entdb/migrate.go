package entdb

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent"
)

func RunMigration(ctx context.Context, client *ent.Client) error {
	log.Info().Msg("🌿 Updating Ent schema migration...")

	return client.Schema.Create(ctx)
}
