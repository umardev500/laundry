package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/pkg/errorsx"
	"github.com/umardev500/laundry/pkg/types"
)

type Tenant struct {
	ID        uuid.UUID
	Name      string
	Phone     string
	Email     string
	Status    types.TenantStatus
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
	t.Status = types.TenantStatusDeleted
	t.DeletedAt = &now
}

// IsDeleted checks whether a tenant has been soft-deleted.
func (t *Tenant) IsDeleted() bool {
	return t.DeletedAt != nil
}

// SetStatus updates the tenant's status with validation.
func (t *Tenant) SetStatus(status types.TenantStatus) error {
	if t.Status == status.Normalize() {
		return types.ErrStatusUnchanged
	}

	if !t.Status.CanTransitionTo(status) {
		allowedStatuses := t.Status.AllowedNextStatuses()
		return errorsx.NewErrInvalidStatusTransition(
			string(t.Status),
			string(status.Normalize()),
			allowedStatuses,
		)
	}

	t.Status = status.Normalize()
	t.UpdatedAt = time.Now().UTC()
	return nil
}

// BelongsToTenant checks whether the service belongs to the tenant in context.
func (s *Tenant) BelongsToTenant(ctx *appctx.Context) bool {
	if ctx.Scope() == appctx.ScopeTenant {
		return ctx.TenantID() != nil && s.ID == *ctx.TenantID()
	}
	return true
}
