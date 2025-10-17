package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/pkg/errorsx"
	"github.com/umardev500/laundry/pkg/types"
)

type Machine struct {
	ID            uuid.UUID
	TenantID      uuid.UUID
	MachineTypeID *uuid.UUID
	Name          string
	Description   string
	Status        types.MachineStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}

func (m *Machine) Update(name, description string) {
	if name != "" {
		m.Name = name
	}
	if description != "" {
		m.Description = description
	}
}

func (m *Machine) SoftDelete() {
	now := time.Now().UTC()
	m.Status = types.MachineStatusOffline // or use Deleted if defined; using Offline as soft-delete marker
	m.DeletedAt = &now
}

func (m *Machine) IsDeleted() bool {
	return m.DeletedAt != nil
}

func (m *Machine) SetStatus(status types.MachineStatus) error {
	if m.Status == status.Normalize() {
		return types.ErrStatusUnchanged
	}

	if !m.Status.CanTransitionTo(status) {
		return errorsx.NewErrInvalidStatusTransition(
			string(m.Status),
			string(status.Normalize()),
			m.Status.AllowedNextStatuses(),
		)
	}

	m.Status = status.Normalize()
	m.UpdatedAt = time.Now().UTC()
	return nil
}

// BelongsToTenant checks whether the machine belongs to the tenant in context.
// Returns true if scope is platform or tenant IDs match.
func (r *Machine) BelongsToTenant(ctx *appctx.Context) bool {
	if ctx.Scope() == appctx.ScopeTenant {
		return r.TenantID != uuid.Nil && ctx.TenantID() != nil && r.TenantID == *ctx.TenantID()
	}
	return true
}
