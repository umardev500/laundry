package domain

import "errors"

var (
	// ErrUserDeleted is returned when attempting to update a soft-deleted user.
	ErrUserDeleted = errors.New("cannot update deleted user")

	// ErrUserSuspended is returned when the user account is suspended.
	ErrUserSuspended = errors.New("user account is suspended")

	// ErrUserNotFound is returned when a user cannot be found.
	ErrUserNotFound = errors.New("user not found")

	ErrUserAlreadyExists = errors.New("user already exists")

	ErrUnauthorizedUserAccess = errors.New("unauthorized access to user")
)
