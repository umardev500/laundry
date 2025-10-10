package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
)

type ServiceCategory struct {
	ID          uuid.UUID
	TenantID    uuid.UUID
	Name        string
	Description string
	DeletedAt   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func New(id, tenantID uuid.UUID, name, description string) *ServiceCategory {
	return &ServiceCategory{
		ID:          id,
		TenantID:    tenantID,
		Name:        name,
		Description: description,
	}
}

// BelongsToTenant checks whether the service belongs to the tenant in context.
func (s *ServiceCategory) BelongsToTenant(ctx *appctx.Context) bool {
	if ctx.Scope() == appctx.ScopeTenant {
		return ctx.TenantID() != nil && s.TenantID == *ctx.TenantID()
	}
	return true
}

func (s *ServiceCategory) Update(name, description string) {
	s.Name = name
	s.Description = description
}

func (s *ServiceCategory) SoftDelete() {
	now := time.Now()
	s.DeletedAt = &now
}

func (s *ServiceCategory) IsDeleted() bool {
	return s.DeletedAt != nil
}
