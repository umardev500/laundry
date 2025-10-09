package seeder

import (
	"context"

	"github.com/rs/zerolog/log"
)

func RunAll(ctx context.Context, seeders []Seeder) error {
	for _, s := range seeders {
		log.Info().Msgf("🌱 Running seeder: %T", s)
		if err := s.Seed(ctx); err != nil {
			log.Fatal().Err(err).Msgf("Failed to run seeder: %T", s)
			return err
		}
	}

	log.Info().Msg("✅ All seeders have been run")
	return nil
}
