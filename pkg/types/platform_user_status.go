package types

import (
	"slices"
	"strings"
)

// PlatformUserStatus represents the status of a PlatformUser.
type PlatformUserStatus string

const (
	PlatformUserStatusActive    PlatformUserStatus = "ACTIVE"
	PlatformUserStatusSuspended PlatformUserStatus = "SUSPENDED"
	PlatformUserStatusDeleted   PlatformUserStatus = "DELETED"
)

// AllowedPlatformUserTransitions defines which PlatformUser statuses can transition to which.
var AllowedPlatformUserTransitions = map[PlatformUserStatus][]PlatformUserStatus{
	PlatformUserStatusActive:    {PlatformUserStatusSuspended, PlatformUserStatusDeleted},
	PlatformUserStatusSuspended: {PlatformUserStatusActive, PlatformUserStatusDeleted},
	PlatformUserStatusDeleted:   {}, // terminal state
}

func (s PlatformUserStatus) CanTransitionTo(next PlatformUserStatus) bool {
	nextNormalize := next.Normalize()
	allowedNext, ok := AllowedPlatformUserTransitions[s]
	if !ok {
		return false
	}

	return slices.Contains(allowedNext, nextNormalize)
}

func (s PlatformUserStatus) AllowedNextStatuses() []PlatformUserStatus {
	return AllowedPlatformUserTransitions[s]
}

func (e PlatformUserStatus) Normalize() PlatformUserStatus {
	return PlatformUserStatus(strings.ToUpper(string(e)))
}
