package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

type Permission struct {
	ID          uuid.UUID
	Name        string
	DisplayName string
	Description string
	Status      types.Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// --- Domain Behaviors ---
func (p *Permission) Update(name, displayName, description string) {
	if name != "" {
		p.Name = name
	}
	if displayName != "" {
		p.DisplayName = displayName
	}
	if description != "" {
		p.Description = description
	}
}

func (p *Permission) SoftDelete() {
	now := time.Now().UTC()
	p.Status = types.StatusDeleted
	p.DeletedAt = &now
}

func (p *Permission) IsDeleted() bool {
	return p.DeletedAt != nil
}

func (p *Permission) SetStatus(status types.Status) error {
	if p.Status == status {
		return types.ErrStatusUnchanged
	}
	if p.IsDeleted() && status != types.StatusDeleted {
		return types.ErrInvalidStatusTransition
	}
	p.Status = status
	p.UpdatedAt = time.Now().UTC()
	return nil
}
