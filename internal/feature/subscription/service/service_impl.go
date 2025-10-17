package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/subscription/contract"
	"github.com/umardev500/laundry/internal/feature/subscription/domain"
	"github.com/umardev500/laundry/internal/feature/subscription/query"
	"github.com/umardev500/laundry/internal/feature/subscription/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type subscriptionService struct {
	repo repository.Repository
}

// NewSubscriptionService creates a new Subscription service.
func NewSubscriptionService(repo repository.Repository) contract.Service {
	return &subscriptionService{
		repo: repo,
	}
}

// Create registers a new subscription.
func (s *subscriptionService) Create(ctx *appctx.Context, sub *domain.Subscription) (*domain.Subscription, error) {
	existing, err := s.repo.FindActiveByTenantID(ctx, sub.TenantID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if existing != nil {
		return nil, domain.ErrSubscriptionAlreadyExists
	}

	if sub.ID == uuid.Nil {
		sub.ID = uuid.New()
	}

	return s.repo.Create(ctx, sub)
}

// Update modifies an existing subscription.
func (s *subscriptionService) Update(ctx *appctx.Context, sub *domain.Subscription) (*domain.Subscription, error) {
	existing, err := s.findExisting(ctx, sub.ID)
	if err != nil {
		return nil, err
	}

	existing.Update(sub)

	return s.repo.Update(ctx, existing)
}

// UpdateStatus updates the status of a subscription.
func (s *subscriptionService) UpdateStatus(ctx *appctx.Context, sub *domain.Subscription) (*domain.Subscription, error) {
	existing, err := s.findExisting(ctx, sub.ID)
	if err != nil {
		return nil, err
	}

	activeExisting, err := s.repo.FindActiveByTenantID(ctx, existing.TenantID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if activeExisting != nil {
		return nil, domain.ErrSubscriptionAlreadyExists
	}

	if err := existing.SetStatus(sub.Status); err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, existing)
}

// Delete performs a soft delete (marks as deleted).
func (s *subscriptionService) Delete(ctx *appctx.Context, id uuid.UUID) error {
	sub, err := s.findExisting(ctx, id)
	if err != nil {
		return err
	}

	sub.Delete() // domain method to update status to DELETED

	_, err = s.repo.Update(ctx, sub)
	return err
}

// Purge permanently removes a subscription from the database.
func (s *subscriptionService) Purge(ctx *appctx.Context, id uuid.UUID) error {
	sub, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, sub.ID)
}

// Restore reactivates a deleted subscription.
func (s *subscriptionService) Restore(ctx *appctx.Context, id uuid.UUID) (*domain.Subscription, error) {
	sub, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return nil, err
	}

	sub.Restore()

	return s.repo.Update(ctx, sub)
}

// FindByID retrieves a subscription by its ID.
func (s *subscriptionService) FindByID(ctx *appctx.Context, id uuid.UUID, q *query.FindSubscriptionByIDQuery) (*domain.Subscription, error) {
	sub, err := s.repo.FindByID(ctx, id, q)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrSubscriptionNotFound
		}
		return nil, err
	}
	if !q.IncludeDeleted && sub.IsDeleted() {
		return nil, domain.ErrSubscriptionDeleted
	}
	return sub, nil
}

// List returns paginated subscriptions.
func (s *subscriptionService) List(ctx *appctx.Context, q *query.ListSubscriptionQuery) (*pagination.PageData[domain.Subscription], error) {
	return s.repo.List(ctx, q)
}

// -------------------------
// Helpers
// -------------------------

// findExisting fetches a subscription that must exist and not be deleted.
func (s *subscriptionService) findExisting(ctx *appctx.Context, id uuid.UUID) (*domain.Subscription, error) {
	q := &query.FindSubscriptionByIDQuery{IncludeDeleted: false}
	sub, err := s.repo.FindByID(ctx, id, q)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrSubscriptionNotFound
		}
		return nil, err
	}
	if sub.IsDeleted() {
		return nil, domain.ErrSubscriptionDeleted
	}
	return sub, nil
}

// findAllowDeleted fetches a subscription that may be deleted (used for purge).
func (s *subscriptionService) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.Subscription, error) {
	q := &query.FindSubscriptionByIDQuery{IncludeDeleted: true}
	sub, err := s.repo.FindByID(ctx, id, q)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrSubscriptionNotFound
		}
		return nil, err
	}
	return sub, nil
}
