package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
)

type ServiceUnit struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	Name      string
	Symbol    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// SoftDelete marks the service unit as deleted without removing the record.
func (s *ServiceUnit) SoftDelete() {
	now := time.Now().UTC()
	s.DeletedAt = &now
}

func (s *ServiceUnit) IsDeleted() bool {
	return s.DeletedAt != nil
}

// Update applies updated data.
func (s *ServiceUnit) Update(name, symbol string) {
	if name != "" {
		s.Name = name
	}
	if symbol != "" {
		s.Symbol = symbol
	}
}

// BelongsToTenant checks whether the unit belongs to the current tenant context.
func (s *ServiceUnit) BelongsToTenant(ctx *appctx.Context) bool {
	if ctx.Scope() == appctx.ScopeTenant {
		return s.TenantID == *ctx.TenantID()
	}
	return true
}
