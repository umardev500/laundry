package domain

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusActive    Status = "active"
	StatusDeleted   Status = "deleted"
	StatusSuspended Status = "suspended"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Status    Status
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
	u.Status = StatusDeleted
	u.DeletedAt = &now
}

// IsDeleted checks whether a user has been soft-deleted.
func (u *User) IsDeleted() bool {
	return u.DeletedAt != nil
}

// SetStatus updates the user's status with validation.
func (u *User) SetStatus(status Status) error {
	if u.Status == status {
		return ErrStatusUnchanged
	}

	// Optional rule: prevent reactivating deleted users.
	if u.IsDeleted() && status != StatusDeleted {
		return ErrInvalidStatusTransition
	}

	u.Status = status
	u.UpdatedAt = time.Now().UTC()
	return nil
}
