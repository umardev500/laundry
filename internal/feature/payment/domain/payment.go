package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/pkg/types"
)

type Payment struct {
	ID              uuid.UUID
	TenantID        *uuid.UUID
	RefID           uuid.UUID
	RefType         types.PaymentType
	PaymentMethodID uuid.UUID
	Amount          float64
	ReceivedAmount  *float64
	ChangeAmount    *float64
	Notes           string
	Status          types.PaymentStatus
	PaidAt          *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

// -------------------------
// Initialization / Defaults
// -------------------------

func (p *Payment) InitDefaults() {
	if p.Status == "" {
		p.Status = types.PaymentStatusPending
	}
}

func (p *Payment) Create() {
	p.InitDefaults()
}

// -------------------------
// Validation
// -------------------------

func (p *Payment) Validate() error {
	if p.RefType != types.PaymentTypeOrder && p.RefType != types.PaymentTypeSubscription {
		return types.ErrInvalidPaymentType
	}

	if p.Status != types.PaymentStatusPending &&
		p.Status != types.PaymentStatusPaid &&
		p.Status != types.PaymentStatusFailed {
		return types.ErrInvalidPaymentStatus
	}

	if p.PaidAt != nil && p.Status != types.PaymentStatusPaid {
		return types.ErrInvalidPaymentStatus
	}

	return nil
}

// -------------------------
// Actions / Mutations
// -------------------------

func (p *Payment) CompleteCashPayment(received float64) error {
	if p.Status != types.PaymentStatusPending {
		return types.ErrOnlyPendingPayments
	}

	if received < p.Amount {
		return types.ErrInsufficientPayment
	}

	now := time.Now()
	p.ReceivedAmount = &received
	change := received - p.Amount
	p.ChangeAmount = &change
	p.Status = types.PaymentStatusPaid
	p.PaidAt = &now
	p.UpdatedAt = now

	return nil
}

// Update updates mutable fields of a payment
func (p *Payment) Update(
	paymentMethodID uuid.UUID,
	amount float64,
	notes string,
	receivedAmount *float64,
) error {
	if p.IsDeleted() {
		return ErrPaymentDeleted
	}

	p.PaymentMethodID = paymentMethodID
	p.Amount = amount
	p.Notes = notes
	p.UpdatedAt = time.Now()

	if p.ReceivedAmount != nil {
		p.ReceivedAmount = receivedAmount
		change := *receivedAmount - amount
		p.ChangeAmount = &change

		// If payment was pending, mark as paid
		if p.Status == types.PaymentStatusPending {
			now := time.Now()
			p.Status = types.PaymentStatusPaid
			p.PaidAt = &now
		}
	}

	return p.Validate()
}

// SoftDelete marks the payment as deleted
func (p *Payment) SoftDelete() {
	now := time.Now()
	p.DeletedAt = &now
}

// -------------------------
// Helpers
// -------------------------

// IsDeleted returns true if the payment is deleted.
func (p *Payment) IsDeleted() bool {
	return p.DeletedAt != nil
}

// BelongsToTenant checks whether the service belongs to the tenant in context.
func (s *Payment) BelongsToTenant(ctx *appctx.Context) bool {
	if ctx.Scope() == appctx.ScopeTenant {
		return ctx.TenantID() != nil && s.TenantID == ctx.TenantID()
	}
	return true
}
