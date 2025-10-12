package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/contract"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/domain"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/query"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type serviceImpl struct {
	repo repository.Repository
}

// NewService returns a new PaymentMethod service
func NewService(repo repository.Repository) contract.Service {
	return &serviceImpl{
		repo: repo,
	}
}

// Create adds a new payment method
func (s *serviceImpl) Create(ctx *appctx.Context, pm *domain.PaymentMethod) (*domain.PaymentMethod, error) {
	existing, err := s.repo.FindByName(ctx, pm.Name)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrPaymentMethodAlreadyExists
	}
	return s.repo.Create(ctx, pm)
}

// Update modifies an existing payment method
func (s *serviceImpl) Update(ctx *appctx.Context, pm *domain.PaymentMethod) (*domain.PaymentMethod, error) {
	existing, err := s.findExisting(ctx, pm.ID)
	if err != nil {
		return nil, err
	}

	existing.Update(pm.Name, pm.Type, pm.Description)
	return s.repo.Update(ctx, existing)
}

// Delete performs soft delete
func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	pm, err := s.findExisting(ctx, id)
	if err != nil {
		return err
	}

	pm.SoftDelete()
	_, err = s.repo.Update(ctx, pm)
	return err
}

// Purge permanently deletes a payment method
func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	pm, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, pm.ID)
}

// GetByID returns a payment method by ID
func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.PaymentMethod, error) {
	return s.findExisting(ctx, id)
}

// GetByName returns a payment method by name
func (s *serviceImpl) GetByName(ctx *appctx.Context, name string) (*domain.PaymentMethod, error) {
	return s.repo.FindByName(ctx, name)
}

// List returns paginated payment methods
func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListPaymentMethodQuery) (*pagination.PageData[domain.PaymentMethod], error) {
	return s.repo.List(ctx, q)
}

// ----------------------------
// Helper methods
// ----------------------------

// findExisting ensures a payment method exists and is not soft-deleted
func (s *serviceImpl) findExisting(ctx *appctx.Context, id uuid.UUID) (*domain.PaymentMethod, error) {
	pm, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrPaymentMethodNotFound
		}
		return nil, err
	}

	if pm.IsDeleted() {
		return nil, domain.ErrPaymentMethodDeleted
	}

	return pm, nil
}

// findAllowDeleted fetches a payment method regardless of deleted status
func (s *serviceImpl) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.PaymentMethod, error) {
	pm, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrPaymentMethodNotFound
		}
		return nil, err
	}

	return pm, nil
}
