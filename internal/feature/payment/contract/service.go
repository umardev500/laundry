package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/payment/domain"
)

// Service defines the business logic for payments.
type Service interface {
	// Create a new payment
	Create(ctx *appctx.Context, payment *domain.Payment) (*domain.Payment, error)

	// Update an existing payment
	Update(ctx *appctx.Context, payment *domain.Payment) (*domain.Payment, error)

	// GetByID retrieves a payment by its ID
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Payment, error)

	// Delete a payment by its ID
	Delete(ctx *appctx.Context, id uuid.UUID) error

	// MarkPaid completes a payment (cash or other methods)
	MarkPaid(ctx *appctx.Context, id uuid.UUID, receivedAmount float64) (*domain.Payment, error)

	// Purge a payment by its ID
	Purge(ctx *appctx.Context, id uuid.UUID) error
}
