package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/errors"
	"github.com/umardev500/laundry/pkg/types"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Status    types.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Update updates the email and password of a user.
func (u *User) Update(email, password string) {
	if email != "" {
		u.Email = email
	}
	if password != "" {
		u.Password = password
	}
}

// SoftDelete marks a user as deleted without removing the record.
func (u *User) SoftDelete() {
	now := time.Now().UTC()
	u.Status = types.UserStatusDeleted
	u.DeletedAt = &now
}

// IsDeleted checks whether a user has been soft-deleted.
func (u *User) IsDeleted() bool {
	return u.DeletedAt != nil
}

// SetStatus updates the user's status with validation.
func (u *User) SetStatus(status types.UserStatus) error {
	status = status.Normalize()
	if u.Status == status {
		return types.ErrStatusUnchanged
	}

	// Optional rule: prevent reactivating deleted users.
	if !u.Status.CanTransitionTo(status) {
		return errors.NewErrInvalidStatusTransition(
			string(u.Status),
			string(status),
			u.Status.AllowedNextStatuses(),
		)
	}

	u.Status = status
	u.UpdatedAt = time.Now().UTC()
	return nil
}
