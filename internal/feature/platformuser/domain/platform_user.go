package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/errorsx"
	"github.com/umardev500/laundry/pkg/types"
)

type PlatformUser struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Status    types.PlatformUserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// SoftDelete marks a user as deleted without removing the record.
func (u *PlatformUser) SoftDelete() {
	now := time.Now().UTC()
	u.Status = types.PlatformUserStatusDeleted
	u.DeletedAt = &now
}

func (p *PlatformUser) IsDeleted() bool {
	return p.DeletedAt != nil
}

func (p *PlatformUser) SetStatus(status types.PlatformUserStatus) error {
	status = status.Normalize()

	if p.Status == status {
		return types.ErrStatusUnchanged
	}

	if !p.Status.CanTransitionTo(status) {
		return errorsx.NewErrInvalidStatusTransition(
			string(p.Status),
			string(status),
			p.Status.AllowedNextStatuses(),
		)
	}

	p.Status = status
	p.UpdatedAt = time.Now().UTC()
	return nil
}
