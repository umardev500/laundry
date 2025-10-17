package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/errorsx"
	"github.com/umardev500/laundry/pkg/types"
)

// Subscription represents a user's subscription to a plan.
type Subscription struct {
	ID        uuid.UUID
	TenantID  uuid.UUID
	PlanID    uuid.UUID
	Status    types.SubscriptionStatus
	StartDate *time.Time
	EndDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// NewSubscription creates a new subscription with default Active status.
func NewSubscription(tenantID, planID uuid.UUID, startDate, endDate *time.Time) *Subscription {
	return &Subscription{
		ID:        uuid.New(),
		TenantID:  tenantID,
		PlanID:    planID,
		Status:    types.SubscriptionStatusPending,
		StartDate: startDate,
		EndDate:   endDate,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

// Restore undeletes a soft-deleted subscription and optionally restores it to Active status.
func (s *Subscription) Restore() error {
	if !s.IsDeleted() {
		// Subscription is not deleted, nothing to do
		return ErrSubscriptionNotDeleted
	}

	// Clear the deletion timestamp
	s.DeletedAt = nil

	return nil
}

// Update applies non-zero or changed fields from another Subscription to this one.
func (s *Subscription) Update(in *Subscription) {
	if in == nil {
		return
	}

	// PlanID can change (e.g., upgrade/downgrade)
	if in.PlanID != uuid.Nil && in.PlanID != s.PlanID {
		s.PlanID = in.PlanID
	}

	// Status can change via normal business rules
	if in.Status != "" && in.Status != s.Status {
		_ = s.SetStatus(in.Status) // ignore invalid transitions here; service should validate
	}

	// StartDate update (if provided and different)
	if in.StartDate == nil && !in.StartDate.Equal(*s.StartDate) {
		s.StartDate = in.StartDate
	}

	// EndDate update (if provided and different)
	if in.EndDate != nil {
		if s.EndDate == nil || !in.EndDate.Equal(*s.EndDate) {
			s.EndDate = in.EndDate
		}
	}
}

// SetStatus updates the subscription status if the transition is allowed.
func (s *Subscription) SetStatus(newStatus types.SubscriptionStatus) error {
	newStatus = newStatus.Normalize()

	// Check allowed transition
	if !s.Status.CanTransitionTo(newStatus) {
		return errorsx.NewErrInvalidStatusTransition(
			string(s.Status),
			string(newStatus),
			s.Status.AllowedNextStatuses(),
		)
	}
	s.Status = newStatus
	return nil
}

// Convenience methods
func (s *Subscription) Activate() error {
	return s.SetStatus(types.SubscriptionStatusActive)
}
func (s *Subscription) Cancel() error {
	return s.SetStatus(types.SubscriptionStatusCanceled)
}
func (s *Subscription) Expire() error {
	return s.SetStatus(types.SubscriptionStatusExpired)
}
func (s *Subscription) Suspend() error {
	return s.SetStatus(types.SubscriptionStatusSuspended)
}
func (s *Subscription) Delete() error {
	if s.IsDeleted() {
		return ErrSubscriptionDeleted
	}

	s.Status = types.SubscriptionStatusDeleted
	now := time.Now().UTC()
	s.DeletedAt = &now
	return nil
}

// --- Helper methods ---

// IsDeleted returns true if the subscription has been soft-deleted.
func (s *Subscription) IsDeleted() bool {
	return s.DeletedAt != nil
}

// IsActive returns true if the subscription is active and not deleted.
func (s *Subscription) IsActive() bool {
	return s.Status == types.SubscriptionStatusActive && !s.IsDeleted()
}
