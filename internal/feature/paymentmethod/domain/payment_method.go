package domain

import (
	"time"

	"github.com/google/uuid"
)

// PaymentMethod represents a payment method entity in the domain
type PaymentMethod struct {
	ID          uuid.UUID
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// InitDefaults sets default values for a new PaymentMethod
func (p *PaymentMethod) InitDefaults() {
	now := time.Now()
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	if p.UpdatedAt.IsZero() {
		p.UpdatedAt = now
	}
}

// Update updates mutable fields
func (p *PaymentMethod) Update(name string, description *string) {
	p.Name = name
	p.Description = description
	p.UpdatedAt = time.Now()
}

// SoftDelete marks the payment method as deleted
func (p *PaymentMethod) SoftDelete() {
	now := time.Now()
	p.DeletedAt = &now
	p.UpdatedAt = now
}

// IsDeleted returns true if the payment method is soft-deleted
func (p *PaymentMethod) IsDeleted() bool {
	return p.DeletedAt != nil
}
