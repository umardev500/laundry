package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

type Feature struct {
	ID          uuid.UUID
	Name        string
	Description string
	Status      types.Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (f *Feature) SoftDelete() {
	now := time.Now().UTC()
	f.DeletedAt = &now
}

func (f *Feature) IsDeleted() bool {
	return f.DeletedAt != nil
}

// Update modifies the feature name and/or description.
func (f *Feature) Update(name, description string) {
	if name != "" {
		f.Name = name
	}
	f.Description = description
	f.UpdatedAt = time.Now().UTC()
}

// SetStatus updates the feature's status with validation.
func (f *Feature) SetStatus(status types.Status) error {
	if f.Status == status {
		return types.ErrStatusUnchanged
	}

	// Optional: prevent reactivating deleted features
	if f.IsDeleted() && status != types.StatusDeleted {
		return types.ErrInvalidStatusTransition
	}

	f.Status = status
	f.UpdatedAt = time.Now().UTC()
	return nil
}
