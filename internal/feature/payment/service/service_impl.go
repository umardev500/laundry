package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/payment/contract"
	"github.com/umardev500/laundry/internal/feature/payment/domain"
	"github.com/umardev500/laundry/internal/feature/payment/repository"

	paymentMethodContract "github.com/umardev500/laundry/internal/feature/paymentmethod/contract"
)

// PaymentServiceImpl implements PaymentService
type PaymentServiceImpl struct {
	repo                 repository.Repository
	paymentMethodService paymentMethodContract.Service
}

// NewPaymentService creates a new PaymentService
func NewPaymentService(
	repo repository.Repository,
	paymentMethodService paymentMethodContract.Service,
) contract.Service {
	return &PaymentServiceImpl{
		repo:                 repo,
		paymentMethodService: paymentMethodService,
	}
}

// Create a new payment
func (s *PaymentServiceImpl) Create(ctx *appctx.Context, p *domain.Payment) (*domain.Payment, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}

	// Check for payment method
	_, err := s.paymentMethodService.GetByID(ctx, p.PaymentMethodID)
	if err != nil {
		return nil, err
	}

	// Create a new payment
	p.Create()

	return s.repo.Create(ctx, p)
}

// Update an existing payment
func (s *PaymentServiceImpl) Update(ctx *appctx.Context, p *domain.Payment) (*domain.Payment, error) {
	existing, err := s.findExisting(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	// Update fields
	err = existing.Update(
		p.PaymentMethodID,
		p.Amount,
		p.Notes,
		p.ReceivedAmount,
	)

	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, existing)
}

// UpdateStatus implements contract.Service.
func (s *PaymentServiceImpl) UpdateStatus(ctx *appctx.Context, payment *domain.Payment) (*domain.Payment, error) {
	payment, err := s.findExisting(ctx, payment.ID)
	if err != nil {
		return nil, err
	}

	// Update status
	err = payment.UpdateStatus(payment.Status)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, payment)
}

// GetByID retrieves a payment by its ID
func (s *PaymentServiceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Payment, error) {
	return s.findExisting(ctx, id)
}

// Delete a payment by its ID (soft delete)
func (s *PaymentServiceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	existing, err := s.findExisting(ctx, id)
	if err != nil {
		return err
	}

	existing.SoftDelete()

	_, err = s.repo.Update(ctx, existing)
	return err
}

// Purge a payment by its ID (hard delete)
func (s *PaymentServiceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	existing, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, existing.ID)
}

// MarkPaid marks a payment as paid (cash or other method)
func (s *PaymentServiceImpl) MarkPaid(ctx *appctx.Context, id uuid.UUID, receivedAmount float64) (*domain.Payment, error) {
	p, err := s.findExisting(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := p.CompleteCashPayment(receivedAmount); err != nil {
		return nil, err
	}

	p.UpdatedAt = time.Now()
	return s.repo.Update(ctx, p)
}

// -----------------------
// Helper methods
// -----------------------

// findExisting ensures the payment exists, is not soft-deleted, and belongs to tenant
func (s *PaymentServiceImpl) findExisting(ctx *appctx.Context, id uuid.UUID) (*domain.Payment, error) {
	p, err := s.repo.FindById(ctx, id)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, domain.ErrPaymentNotFound
		}
		return nil, err
	}

	if p.IsDeleted() {
		return nil, domain.ErrPaymentDeleted
	}

	if !p.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedPaymentAccess
	}

	return p, nil
}

// findAllowDeleted fetches a payment regardless of deleted status but checks tenant ownership
func (s *PaymentServiceImpl) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.Payment, error) {
	p, err := s.repo.FindById(ctx, id)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, domain.ErrPaymentNotFound
		}
		return nil, err
	}

	if !p.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedPaymentAccess
	}

	return p, nil
}
