package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/pkg/errorsx"
	"github.com/umardev500/laundry/pkg/types"
)

type TenantUser struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	TenantID  uuid.UUID
	Status    types.TenantUserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Create initializes a TenantUser and validates that the tenant matches
// the logged-in user's tenant ID.
func (tu *TenantUser) Create(loggedInTenantID *uuid.UUID) error {
	if loggedInTenantID != nil && tu.TenantID != *loggedInTenantID {
		return ErrTenantIDMismatch
	}

	if tu.ID == uuid.Nil {
		tu.ID = uuid.New()
	}

	tu.Status = types.TenantUserStatusActive
	tu.CreatedAt = time.Now().UTC()
	tu.UpdatedAt = tu.CreatedAt
	tu.DeletedAt = nil

	return nil
}

// SoftDelete marks the tenant user as deleted.
func (tu *TenantUser) SoftDelete() {
	now := time.Now().UTC()
	tu.Status = types.TenantUserStatusDeleted
	tu.DeletedAt = &now
}

// IsDeleted checks if the tenant user is soft-deleted.
func (tu *TenantUser) IsDeleted() bool {
	return tu.DeletedAt != nil
}

// SetStatus safely updates the status.
func (tu *TenantUser) SetStatus(status types.TenantUserStatus) error {
	if tu.Status == status {
		return types.ErrStatusUnchanged
	}

	if !tu.Status.CanTransitionTo(status) {
		return errorsx.NewErrInvalidStatusTransition(
			string(tu.Status),
			string(status),
			tu.Status.AllowedNextStatuses(),
		)
	}

	tu.Status = status
	tu.UpdatedAt = time.Now().UTC()
	return nil
}

// BelongsToTenant checks whether the machine belongs to the tenant in context.
// Returns true if scope is platform or tenant IDs match.
func (r *TenantUser) BelongsToTenant(ctx *appctx.Context) bool {
	if ctx.Scope() == appctx.ScopeTenant {
		return r.TenantID != uuid.Nil && ctx.TenantID() != nil && r.TenantID == *ctx.TenantID()
	}
	return true
}
