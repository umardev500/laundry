package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

type Tenant struct {
	ID        uuid.UUID
	Name      string
	Phone     string
	Email     string
	Status    types.Status
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Update updates the name and email of a tenant.
func (t *Tenant) Update(name, email string) {
	if name != "" {
		t.Name = name
	}
	if email != "" {
		t.Email = email
	}
}

// SoftDelete marks a tenant as deleted without removing the record.
func (t *Tenant) SoftDelete() {
	now := time.Now().UTC()
	t.Status = types.StatusDeleted
	t.DeletedAt = &now
}

// IsDeleted checks whether a tenant has been soft-deleted.
func (t *Tenant) IsDeleted() bool {
	return t.DeletedAt != nil
}

// SetStatus updates the tenant's status with validation.
func (t *Tenant) SetStatus(status types.Status) error {
	if t.Status == status {
		return ErrStatusUnchanged
	}

	if t.IsDeleted() && status != types.StatusDeleted {
		return ErrInvalidStatusTransition
	}

	t.Status = status
	t.UpdatedAt = time.Now().UTC()
	return nil
}
