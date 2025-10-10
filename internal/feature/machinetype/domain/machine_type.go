package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
)

type MachineType struct {
	ID          uuid.UUID
	TenantID    uuid.UUID
	Name        string
	Description string
	Capacity    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// Update applies new values.
func (m *MachineType) Update(name, description string, capacity int) {
	if name != "" {
		m.Name = name
	}
	if description != "" {
		m.Description = description
	}
	if capacity > 0 {
		m.Capacity = capacity
	}
}

func (m *MachineType) SoftDelete() {
	now := time.Now().UTC()
	m.DeletedAt = &now
}

func (m *MachineType) IsDeleted() bool {
	return m.DeletedAt != nil
}

// BelongsToTenant checks whether the machine type belongs to the tenant in context.
// Returns true if scope is platform or tenant IDs match.
func (r *MachineType) BelongsToTenant(ctx *appctx.Context) bool {
	if ctx.Scope() == appctx.ScopeTenant {
		return r.TenantID != uuid.Nil && ctx.TenantID() != nil && r.TenantID == *ctx.TenantID()
	}
	return true
}
