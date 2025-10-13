package seeder

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/paymentmethod"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/types"
)

type PaymentMethodSeeder struct {
	client *entdb.Client
}

func NewPaymentMethodSeeder(client *entdb.Client) *PaymentMethodSeeder {
	return &PaymentMethodSeeder{client: client}
}

func (s *PaymentMethodSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding payment methods...")

	conn := s.client.GetConn(ctx)

	methods := []struct {
		ID          uuid.UUID
		Name        string
		Description string
		Type        types.PaymentMethod
	}{
		{
			ID:          uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Name:        "Cash",
			Description: "Cash payment",
			Type:        types.PaymentMethodCash,
		},
		{
			ID:          uuid.MustParse("22222222-1111-1111-1111-111111111111"),
			Name:        "Card",
			Description: "Credit/Debit Card payment",
			Type:        types.PaymentMethodCard,
		},
		{
			ID:          uuid.MustParse("33333333-1111-1111-1111-111111111111"),
			Name:        "Transfer",
			Description: "Bank Transfer payment",
			Type:        types.PaymentMethodTransfer,
		},
	}

	for _, m := range methods {
		exists, _ := conn.PaymentMethod.Query().
			Where(paymentmethod.IDEQ(m.ID)).
			Exist(ctx)

		if !exists {
			_, err := conn.PaymentMethod.Create().
				SetID(m.ID).
				SetName(m.Name).
				SetNillableDescription(&m.Description).
				SetType(paymentmethod.Type(m.Type)).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to seed payment method %s: %w", m.Name, err)
			}
		}
	}

	return nil
}
