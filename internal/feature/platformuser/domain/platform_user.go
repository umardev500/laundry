package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

type PlatformUser struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Status    types.Status
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// SoftDelete marks a user as deleted without removing the record.
func (u *PlatformUser) SoftDelete() {
	now := time.Now().UTC()
	u.Status = types.StatusDeleted
	u.DeletedAt = &now
}

func (p *PlatformUser) IsDeleted() bool {
	return p.DeletedAt != nil
}

func (p *PlatformUser) SetStatus(status types.Status) error {
	if p.Status == status {
		return ErrStatusUnchanged
	}

	// Optional rule: prevent reactivating deleted users.
	if p.IsDeleted() && status != types.StatusDeleted {
		return ErrInvalidStatusTransition
	}

	p.Status = status
	p.UpdatedAt = time.Now().UTC()
	return nil
}
