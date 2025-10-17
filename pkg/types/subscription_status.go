package types

import (
	"slices"
	"strings"
)

// SubscriptionStatus represents the status of a Subscription.
type SubscriptionStatus string

const (
	SubscriptionStatusPending   SubscriptionStatus = "PENDING"   // Newly created but not yet active (e.g., awaiting payment)
	SubscriptionStatusActive    SubscriptionStatus = "ACTIVE"    // Currently active and billed
	SubscriptionStatusCanceled  SubscriptionStatus = "CANCELED"  // User canceled; may remain active until end_date
	SubscriptionStatusExpired   SubscriptionStatus = "EXPIRED"   // Past end_date; no longer active
	SubscriptionStatusSuspended SubscriptionStatus = "SUSPENDED" // Temporarily disabled (e.g. payment failed)
	SubscriptionStatusDeleted   SubscriptionStatus = "DELETED"   // Permanently removed (soft delete marker)
)

// AllowedSubscriptionTransitions defines valid state transitions for a subscription lifecycle.
// AllowedSubscriptionTransitions defines valid state transitions for a subscription lifecycle.
var AllowedSubscriptionTransitions = map[SubscriptionStatus][]SubscriptionStatus{
	SubscriptionStatusPending: {
		SubscriptionStatusActive,
		SubscriptionStatusCanceled,
	},
	SubscriptionStatusActive: {
		SubscriptionStatusSuspended,
		SubscriptionStatusCanceled,
	},
	SubscriptionStatusSuspended: {
		SubscriptionStatusActive,
		SubscriptionStatusCanceled,
	},
	SubscriptionStatusCanceled: {
		SubscriptionStatusExpired,
	},
	SubscriptionStatusExpired: {}, // expired row is immutable
	SubscriptionStatusDeleted: {}, // deleted row is terminal
}

func (s SubscriptionStatus) CanTransitionTo(next SubscriptionStatus) bool {
	nextNormalize := next.Normalize()
	allowedNext, ok := AllowedSubscriptionTransitions[s]
	if !ok {
		return false
	}

	return slices.Contains(allowedNext, nextNormalize)
}

func (s SubscriptionStatus) AllowedNextStatuses() []SubscriptionStatus {
	return AllowedSubscriptionTransitions[s]
}

func (e SubscriptionStatus) Normalize() SubscriptionStatus {
	return SubscriptionStatus(strings.ToUpper(string(e)))
}
