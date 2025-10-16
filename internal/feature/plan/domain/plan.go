package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// Plan represents the business entity for a subscription plan.
type Plan struct {
	ID              uuid.UUID
	Name            string
	Description     *string
	Price           float64
	BillingInterval types.BillingInterval
	Features        *PlanFeatures
	Active          bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

// NewPlan creates a new Plan with default values.
func NewPlan(name string, description *string, price float64, billingInterval types.BillingInterval, features *PlanFeatures) *Plan {
	return &Plan{
		Name:            name,
		Price:           price,
		BillingInterval: billingInterval,
		Features:        features,
	}
}

// Update updates the Plan with non-nil fields from PlanUpdate.
func (p *Plan) Update(update *Plan) {
	if update.Name != "" {
		p.Name = update.Name
	}
	if update.Description != nil {
		p.Description = update.Description
	}
	if update.Price != 0 {
		p.Price = update.Price
	}
	if update.BillingInterval != "" {
		p.BillingInterval = update.BillingInterval.Normalize()
	}
	if update.Features != nil {
		p.Features = update.Features
	}
}

// IsDeleted returns true if the plan has been soft-deleted.
func (p *Plan) IsDeleted() bool {
	return p.DeletedAt != nil
}

// SoftDelete marks the plan as deleted by setting DeletedAt timestamp.
func (p *Plan) SoftDelete() {
	now := time.Now().UTC()
	p.DeletedAt = &now
}

// Restore restores a soft-deleted plan by clearing the DeletedAt timestamp.
func (p *Plan) Restore() {
	p.DeletedAt = nil
}

// Activate marks the plan as active.
func (p *Plan) Activate() {
	p.Active = true
}

// Deactivate marks the plan as inactive.
func (p *Plan) Deactivate() {
	p.Active = false
}

// IsActive returns true if the plan is active and not deleted.
func (p *Plan) IsActive() bool {
	return p.Active && !p.IsDeleted()
}
