package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/payment/domain"
	"github.com/umardev500/laundry/internal/feature/payment/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	// Create inserts a new payment into the database.
	Create(ctx *appctx.Context, p *domain.Payment) (*domain.Payment, error)

	// Update updates an existing payment.
	Update(ctx *appctx.Context, p *domain.Payment) (*domain.Payment, error)

	// FindById returns a payment by its ID.
	FindById(ctx *appctx.Context, id uuid.UUID) (*domain.Payment, error)

	// Delete deletes a payment by its ID.
	Delete(ctx *appctx.Context, id uuid.UUID) error

	List(ctx *appctx.Context, f *query.ListPaymentQuery) (*pagination.PageData[domain.Payment], error)
}
