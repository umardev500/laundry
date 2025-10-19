package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/address/contract"
	"github.com/umardev500/laundry/internal/feature/address/domain"
	"github.com/umardev500/laundry/internal/feature/address/query"
	"github.com/umardev500/laundry/internal/feature/address/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type addressService struct {
	repo repository.Repository
}

// NewAddressService creates a new Address service.
func NewAddressService(repo repository.Repository) contract.Address {
	return &addressService{
		repo: repo,
	}
}

// Create adds a new address for a user.
func (s *addressService) Create(ctx *appctx.Context, a *domain.Address) (*domain.Address, error) {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}

	// If new address is primary, unmark existing ones
	if a.IsPrimary {
		if err := s.repo.UnsetPrimary(ctx, a.User.ID); err != nil {
			return nil, err
		}
	}

	return s.repo.Create(ctx, a)
}

// List retrieves paginated addresses based on query filters.
func (s *addressService) List(ctx *appctx.Context, q *query.ListAddressQuery) (*pagination.PageData[domain.Address], error) {
	return s.repo.List(ctx, q)
}

// GetByID fetches an address by ID.
func (s *addressService) GetByID(ctx *appctx.Context, id uuid.UUID, q *query.FindAddressByIDQuery) (*domain.Address, error) {
	return s.findExisting(ctx, id, q)
}

// Update modifies an existing address.
func (s *addressService) Update(ctx *appctx.Context, a *domain.Address) (*domain.Address, error) {
	existing, err := s.findExisting(ctx, a.ID, &query.FindAddressByIDQuery{
		WithUser:     true,
		WithProvince: true,
		WithRegency:  true,
		WithDistrict: true,
		WithVillage:  true,
	})
	if err != nil {
		return nil, err
	}

	isHavePrimary := existing.IsPrimary

	existing.Update(a)

	// Handle primary address change
	if a.IsPrimary && !isHavePrimary {
		if err := s.repo.UnsetPrimary(ctx, existing.User.ID); err != nil {
			return nil, err
		}
		existing.IsPrimary = true
	}

	return s.repo.Update(ctx, existing)
}

// Delete performs a soft delete on an address.
func (s *addressService) Delete(ctx *appctx.Context, id uuid.UUID) error {
	a, err := s.findExisting(ctx, id, &query.FindAddressByIDQuery{
		WithUser:     true,
		WithProvince: true,
		WithRegency:  true,
		WithDistrict: true,
		WithVillage:  true,
	})
	if err != nil {
		return err
	}

	a.SoftDelete()

	_, err = s.repo.Update(ctx, a)
	return err
}

// Purge permanently deletes an address.
func (s *addressService) Purge(ctx *appctx.Context, id uuid.UUID) error {
	a, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, a.ID)
}

// Restore undeletes a previously deleted address.
func (s *addressService) Restore(ctx *appctx.Context, id uuid.UUID) (*domain.Address, error) {
	a, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return nil, err
	}

	if !a.IsDeleted() {
		return a, nil // already active
	}
	a.Restore()

	return s.repo.Update(ctx, a)
}

// GetPrimaryByUserID fetches the primary address of a user.
func (s *addressService) GetPrimaryByUserID(ctx *appctx.Context, userID uuid.UUID) (*domain.Address, error) {
	addr, err := s.repo.FindPrimaryByUserID(ctx, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.NewAddressError(domain.ErrAddressNotFound)
		}
		return nil, err
	}
	if addr.IsDeleted() {
		return nil, domain.NewAddressError(domain.ErrAddressDeleted)
	}
	return addr, nil
}

// SetPrimary marks the given address as the user's primary.
func (s *addressService) SetPrimary(ctx *appctx.Context, id uuid.UUID, userID uuid.UUID) (*domain.Address, error) {
	addr, err := s.findExisting(ctx, id, &query.FindAddressByIDQuery{
		WithUser:     true,
		WithProvince: true,
		WithRegency:  true,
		WithDistrict: true,
		WithVillage:  true,
	})
	if err != nil {
		return nil, err
	}

	// Unmark others, then set this one
	if err := s.repo.UnsetPrimary(ctx, userID); err != nil {
		return nil, err
	}

	addr.IsPrimary = true
	return s.repo.Update(ctx, addr)
}

// -------------------------
// Helpers
// -------------------------

func (s *addressService) findExisting(ctx *appctx.Context, id uuid.UUID, q *query.FindAddressByIDQuery) (*domain.Address, error) {
	a, err := s.repo.FindByID(ctx, id, q)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.NewAddressError(domain.ErrAddressNotFound)
		}
		return nil, err
	}
	if a.IsDeleted() && !q.IncludeDeleted {
		return nil, domain.NewAddressError(domain.ErrAddressDeleted)
	}
	return a, nil
}

func (s *addressService) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.Address, error) {
	a, err := s.repo.FindByID(ctx, id, nil)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.NewAddressError(domain.ErrAddressNotFound)
		}
		return nil, err
	}
	return a, nil
}
