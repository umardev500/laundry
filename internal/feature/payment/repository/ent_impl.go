package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent/payment"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/payment/domain"
	"github.com/umardev500/laundry/internal/feature/payment/mapper"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

// EntPaymentRepository implements domain.Payment repository using Ent
type EntPaymentRepository struct {
	client *entdb.Client
}

// NewEntPaymentRepository creates a new repository instance
func NewEntPaymentRepository(client *entdb.Client) Repository {
	return &EntPaymentRepository{
		client: client,
	}
}

// Create inserts a new payment
func (r *EntPaymentRepository) Create(ctx *appctx.Context, p *domain.Payment) (*domain.Payment, error) {
	entPayment, err := r.client.Client.Payment.
		Create().
		SetRefID(p.RefID).
		SetRefType(payment.RefType(p.RefType)).
		SetPaymentMethodID(p.PaymentMethodID).
		SetAmount(p.Amount).
		SetStatus(payment.Status(p.Status)).
		SetNotes(p.Notes).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetNillableReceivedAmount(p.ReceivedAmount).
		SetNillableChangeAmount(p.ChangeAmount).
		SetNillablePaidAt(p.PaidAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entPayment), nil
}

// Update modifies an existing payment
func (r *EntPaymentRepository) Update(ctx *appctx.Context, p *domain.Payment) (*domain.Payment, error) {
	entPayment, err := r.client.Client.Payment.
		UpdateOneID(p.ID).
		SetPaymentMethodID(p.PaymentMethodID).
		SetAmount(p.Amount).
		SetStatus(payment.Status(p.Status)).
		SetNotes(p.Notes).
		SetUpdatedAt(time.Now()).
		SetNillableReceivedAmount(p.ReceivedAmount).
		SetNillableChangeAmount(p.ChangeAmount).
		SetNillablePaidAt(p.PaidAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entPayment), nil
}

// FindById returns a payment by its ID
func (r *EntPaymentRepository) FindById(ctx *appctx.Context, id uuid.UUID) (*domain.Payment, error) {
	entPayment, err := r.client.Client.Payment.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entPayment), nil
}

// Delete performs a hard delete
func (r *EntPaymentRepository) Delete(ctx *appctx.Context, id uuid.UUID) error {
	err := r.client.Client.Payment.
		DeleteOneID(id).
		Exec(ctx)
	return err
}
