package domain

import (
	"time"

	"github.com/google/uuid"

	regionDomain "github.com/umardev500/laundry/internal/feature/region/domain"
	userDomain "github.com/umardev500/laundry/internal/feature/user/domain"
)

// Address represents a physical address belonging to a tenant or user.
type Address struct {
	ID        uuid.UUID
	User      *userDomain.User
	Province  *regionDomain.Province
	Regency   *regionDomain.Regency
	District  *regionDomain.District
	Village   *regionDomain.Village
	Street    *string
	IsPrimary bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// NewAddress creates a new Address instance with sensible defaults.
func NewAddress(
	userID uuid.UUID,
	provinceID, regencyID, districtID, villageID string,
	street *string,
	isPrimary *bool,
) (*Address, error) {
	if provinceID == "" || regencyID == "" || districtID == "" || villageID == "" {
		return nil, NewAddressError(ErrInvalidAddressLocation)
	}

	a := &Address{
		User: &userDomain.User{
			ID: userID,
		},
		Province: &regionDomain.Province{
			ID: provinceID,
		},
		Regency: &regionDomain.Regency{
			ID: regencyID,
		},
		District: &regionDomain.District{
			ID: districtID,
		},
		Village: &regionDomain.Village{
			ID: villageID,
		},
		Street: street,
	}

	if isPrimary != nil {
		a.IsPrimary = *isPrimary
	}

	return a, nil
}

// Update applies changes from another Address instance.
// Only non-zero and changed fields will be updated.
func (a *Address) Update(in *Address) {
	if in == nil {
		return
	}

	// --- Province ---
	if in.Province != nil {
		if a.Province == nil || (in.Province.ID != "" && in.Province.ID != a.Province.ID) {
			a.Province = in.Province
		}
	}

	// --- Regency ---
	if in.Regency != nil {
		if a.Regency == nil || (in.Regency.ID != "" && in.Regency.ID != a.Regency.ID) {
			a.Regency = in.Regency
		}
	}

	// --- District ---
	if in.District != nil {
		if a.District == nil || (in.District.ID != "" && in.District.ID != a.District.ID) {
			a.District = in.District
		}
	}

	// --- Village ---
	if in.Village != nil {
		if a.Village == nil || (in.Village.ID != "" && in.Village.ID != a.Village.ID) {
			a.Village = in.Village
		}
	}

	// --- Street ---
	if in.Street != nil {
		if a.Street == nil || *a.Street != *in.Street {
			a.Street = in.Street
		}
	}

	a.IsPrimary = in.IsPrimary
}

// SoftDelete marks the address as deleted.
func (a *Address) SoftDelete() error {
	if a.IsDeleted() {
		return NewAddressError(ErrAddressDeleted)
	}

	now := time.Now().UTC()
	a.DeletedAt = &now
	return nil
}

// Restore undeletes a soft-deleted address.
func (a *Address) Restore() error {
	if !a.IsDeleted() {
		return NewAddressError(ErrAddressNotDeleted)
	}

	a.DeletedAt = nil
	a.UpdatedAt = time.Now().UTC()
	return nil
}

// --- Helper methods ---

// IsDeleted returns true if the address has been soft-deleted.
func (a *Address) IsDeleted() bool {
	return a.DeletedAt != nil
}

// IsActive returns true if the address is not deleted.
func (a *Address) IsActive() bool {
	return !a.IsDeleted()
}
