package types

import (
	"slices"
	"strings"
)

// SubscriptionStatus represents the status of a Subscription.
type SubscriptionStatus string

const (
	SubscriptionStatusActive    SubscriptionStatus = "ACTIVE"    // Currently active and billed
	SubscriptionStatusCanceled  SubscriptionStatus = "CANCELED"  // User canceled; may remain active until end_date
	SubscriptionStatusExpired   SubscriptionStatus = "EXPIRED"   // Past end_date; no longer active
	SubscriptionStatusSuspended SubscriptionStatus = "SUSPENDED" // Temporarily disabled (e.g. payment failed)
	SubscriptionStatusDeleted   SubscriptionStatus = "DELETED"   // Permanently removed (soft delete marker)
)

// AllowedSubscriptionTransitions defines valid state transitions for a subscription lifecycle.
var AllowedSubscriptionTransitions = map[SubscriptionStatus][]SubscriptionStatus{
	SubscriptionStatusActive: {
		SubscriptionStatusSuspended,
		SubscriptionStatusCanceled,
		SubscriptionStatusDeleted,
		SubscriptionStatusActive, // renewal (extend current cycle)
	},
	SubscriptionStatusSuspended: {
		SubscriptionStatusActive, // resume / renew after suspension
		SubscriptionStatusCanceled,
		SubscriptionStatusDeleted,
	},
	SubscriptionStatusCanceled: {
		SubscriptionStatusActive, // renewal or user reactivation
		SubscriptionStatusExpired,
		SubscriptionStatusDeleted,
	},
	SubscriptionStatusExpired: {
		SubscriptionStatusActive, // manual renewal after expiration
	},
	SubscriptionStatusDeleted: {}, // terminal state — cannot renew
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
