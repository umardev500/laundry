package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
)

type Service struct {
	ID                uuid.UUID
	TenantID          uuid.UUID
	ServiceUnitID     *uuid.UUID
	ServiceCategoryID *uuid.UUID
	Name              string
	BasePrice         float64
	Description       string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}

// Update applies changes to mutable fields.
func (s *Service) Update(name string, price float64, description string, unitID, categoryID *uuid.UUID) {
	if name != "" {
		s.Name = name
	}
	// allow price 0.0 as valid; only update if non-negative (you can adjust rules)
	if price >= 0 {
		s.BasePrice = price
	}
	if description != "" {
		s.Description = description
	}
	if unitID != nil {
		s.ServiceUnitID = unitID
	}
	if categoryID != nil {
		s.ServiceCategoryID = categoryID
	}
}

// SoftDelete marks record as deleted.
func (s *Service) SoftDelete() {
	now := time.Now().UTC()
	s.DeletedAt = &now
}

// IsDeleted returns true if deleted_at is set.
func (s *Service) IsDeleted() bool {
	return s.DeletedAt != nil
}

// BelongsToTenant checks whether the service belongs to the tenant in context.
func (s *Service) BelongsToTenant(ctx *appctx.Context) bool {
	if ctx.Scope() == appctx.ScopeTenant {
		return ctx.TenantID() != nil && s.TenantID == *ctx.TenantID()
	}
	return true
}
