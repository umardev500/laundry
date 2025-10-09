package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

type TenantUser struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	TenantID  uuid.UUID
	Status    types.Status
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// SoftDelete marks the tenant user as deleted.
func (tu *TenantUser) SoftDelete() {
	now := time.Now().UTC()
	tu.Status = types.StatusDeleted
	tu.DeletedAt = &now
}

// IsDeleted checks if the tenant user is soft-deleted.
func (tu *TenantUser) IsDeleted() bool {
	return tu.DeletedAt != nil
}

// SetStatus safely updates the status.
func (tu *TenantUser) SetStatus(status types.Status) error {
	if tu.Status == status {
		return types.ErrStatusUnchanged
	}

	if tu.IsDeleted() && status != types.StatusDeleted {
		return types.ErrInvalidStatusTransition
	}

	tu.Status = status
	tu.UpdatedAt = time.Now().UTC()
	return nil
}
