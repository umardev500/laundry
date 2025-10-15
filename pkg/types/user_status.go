package types

import (
	"slices"
	"strings"
)

// UserStatus represents the status of a User.
type UserStatus string

const (
	UserStatusActive    UserStatus = "ACTIVE"
	UserStatusSuspended UserStatus = "SUSPENDED"
	UserStatusDeleted   UserStatus = "DELETED"
)

// AllowedUserTransitions defines which User statuses can transition to which.
var AllowedUserTransitions = map[UserStatus][]UserStatus{
	UserStatusActive:    {UserStatusSuspended, UserStatusDeleted},
	UserStatusSuspended: {UserStatusActive, UserStatusDeleted},
	UserStatusDeleted:   {}, // terminal state
}

func (s UserStatus) CanTransitionTo(next UserStatus) bool {
	nextNormalize := next.Normalize()
	allowedNext, ok := AllowedUserTransitions[s]
	if !ok {
		return false
	}

	return slices.Contains(allowedNext, nextNormalize)
}

func (s UserStatus) AllowedNextStatuses() []UserStatus {
	return AllowedUserTransitions[s]
}

func (e UserStatus) Normalize() UserStatus {
	return UserStatus(strings.ToUpper(string(e)))
}
