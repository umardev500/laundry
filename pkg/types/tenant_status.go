package types

import (
	"slices"
	"strings"
)

// TenantStatus represents the status of a tenant.
type TenantStatus string

const (
	TenantStatusActive    TenantStatus = "ACTIVE"
	TenantStatusSuspended TenantStatus = "SUSPENDED"
	TenantStatusDeleted   TenantStatus = "DELETED"
)

// AllowedTenantTransitions defines which tenant statuses can transition to which.
var AllowedTenantTransitions = map[TenantStatus][]TenantStatus{
	TenantStatusActive:    {TenantStatusSuspended, TenantStatusDeleted},
	TenantStatusSuspended: {TenantStatusActive, TenantStatusDeleted},
	TenantStatusDeleted:   {}, // terminal state
}

func (s TenantStatus) CanTransitionTo(next TenantStatus) bool {
	allowedNext, ok := AllowedTenantTransitions[s]
	if !ok {
		return false
	}
	return slices.Contains(allowedNext, next.Normalize())
}

func (s TenantStatus) AllowedNextStatuses() []TenantStatus {
	return AllowedTenantTransitions[s]
}

func (e TenantStatus) Normalize() TenantStatus {
	return TenantStatus(strings.ToUpper(string(e)))
}
