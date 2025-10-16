package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/plan/contract"
	"github.com/umardev500/laundry/internal/feature/plan/domain"
	"github.com/umardev500/laundry/internal/feature/plan/query"
	"github.com/umardev500/laundry/internal/feature/plan/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type planService struct {
	repo repository.Repository
}

// NewPlanService creates a new Plan service.
func NewPlanService(repo repository.Repository) contract.Plan {
	return &planService{
		repo: repo,
	}
}

// Activate marks a plan as active.
func (s *planService) Activate(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error) {
	p, err := s.findExisting(ctx, id)
	if err != nil {
		return nil, err
	}

	if p.IsActive() {
		return p, nil
	}

	p.Activate()
	return s.repo.Update(ctx, p)
}

// Deactivate marks a plan as inactive.
func (s *planService) Deactivate(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error) {
	p, err := s.findExisting(ctx, id)
	if err != nil {
		return nil, err
	}

	p.Deactivate()

	return s.repo.Update(ctx, p)
}

// Restore clears the soft-deleted flag on a plan.
func (s *planService) Restore(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error) {
	p, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return nil, err
	}

	if !p.IsDeleted() {
		return p, nil // already active
	}
	p.Restore()

	return s.repo.Update(ctx, p)
}

// Create adds a new plan.
func (s *planService) Create(ctx *appctx.Context, p *domain.Plan) (*domain.Plan, error) {
	existing, err := s.repo.FindByName(ctx, p.Name)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if existing != nil && !existing.IsDeleted() {
		return nil, domain.ErrPlanAlreadyExists
	}

	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}

	return s.repo.Create(ctx, p)
}

// List retrieves paginated plans.
func (s *planService) List(ctx *appctx.Context, q *query.ListPlanQuery) (*pagination.PageData[domain.Plan], error) {
	return s.repo.List(ctx, q)
}

// GetByID fetches a plan by ID.
func (s *planService) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error) {
	return s.findExisting(ctx, id)
}

// GetByName fetches a plan by name.
func (s *planService) GetByName(ctx *appctx.Context, name string) (*domain.Plan, error) {
	p, err := s.repo.FindByName(ctx, name)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrPlanNotFound
		}
		return nil, err
	}
	if p.IsDeleted() {
		return nil, domain.ErrPlanDeleted
	}
	return p, nil
}

// Update modifies an existing plan.
func (s *planService) Update(ctx *appctx.Context, p *domain.Plan) (*domain.Plan, error) {
	existing, err := s.findExisting(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	existing.Update(p) // update non-nil fields

	return s.repo.Update(ctx, existing)
}

// Delete performs a soft delete on the plan.
func (s *planService) Delete(ctx *appctx.Context, id uuid.UUID) error {
	p, err := s.findExisting(ctx, id)
	if err != nil {
		return err
	}

	p.SoftDelete()
	_, err = s.repo.Update(ctx, p)
	return err
}

// Purge performs a hard delete on the plan.
func (s *planService) Purge(ctx *appctx.Context, id uuid.UUID) error {
	p, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, p.ID)
}

// -------------------------
// Helpers
// -------------------------

// findExisting fetches a plan that must exist and not be deleted.
func (s *planService) findExisting(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrPlanNotFound
		}
		return nil, err
	}
	if p.IsDeleted() {
		return nil, domain.ErrPlanDeleted
	}
	return p, nil
}

// findAllowDeleted fetches a plan that may be deleted (used for purge).
func (s *planService) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error) {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrPlanNotFound
		}
		return nil, err
	}
	return p, nil
}
