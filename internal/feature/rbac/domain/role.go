package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
)

type Role struct {
	ID          uuid.UUID
	TenantID    *uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// Update updates the name and description of the role.
func (r *Role) Update(name, description string) {
	if name != "" {
		r.Name = name
	}
	if description != "" {
		r.Description = description
	}
	r.UpdatedAt = time.Now().UTC()
}

// SoftDelete marks the role as deleted without removing the record.
func (r *Role) SoftDelete() {
	now := time.Now().UTC()
	r.DeletedAt = &now
	r.UpdatedAt = now
}

// IsDeleted checks whether a role has been soft-deleted.
func (r *Role) IsDeleted() bool {
	return r.DeletedAt != nil
}

// BelongsToTenant checks whether the role belongs to the given tenant.
func (r *Role) BelongsToTenant(ctx *appctx.Context) bool {

	if ctx.Scope() == appctx.ScopeTenant {
		return r.TenantID != nil && ctx.TenantID() != nil && *r.TenantID == *ctx.TenantID()
	}

	return true
}
