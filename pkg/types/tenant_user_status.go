package types

import (
	"slices"
	"strings"
)

// TenantStatus represents the status of a tenant.
type TenantUserStatus string

const (
	TenantUserStatusActive    TenantUserStatus = "ACTIVE"
	TenantUserStatusSuspended TenantUserStatus = "SUSPENDED"
	TenantUserStatusDeleted   TenantUserStatus = "DELETED"
)

// AllowedTenantUserTransitions defines which tenant statuses can transition to which.
var AllowedTenantUserTransitions = map[TenantUserStatus][]TenantUserStatus{
	TenantUserStatusActive:    {TenantUserStatusSuspended, TenantUserStatusDeleted},
	TenantUserStatusSuspended: {TenantUserStatusActive, TenantUserStatusDeleted},
	TenantUserStatusDeleted:   {}, // terminal state
}

func (s TenantUserStatus) CanTransitionTo(next TenantUserStatus) bool {
	allowedNext, ok := AllowedTenantUserTransitions[s]
	if !ok {
		return false
	}
	return slices.Contains(allowedNext, next.Normalize())
}

func (s TenantUserStatus) AllowedNextStatuses() []TenantUserStatus {
	return AllowedTenantUserTransitions[s]
}

func (e TenantUserStatus) Normalize() TenantUserStatus {
	return TenantUserStatus(strings.ToUpper(string(e)))
}
