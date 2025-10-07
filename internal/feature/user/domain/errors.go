package domain

import "errors"

var (
	// ErrUserDeleted is returned when attempting to update a soft-deleted user.
	ErrUserDeleted = errors.New("cannot update deleted user")

	// ErrUserSuspended is returned when the user account is suspended.
	ErrUserSuspended = errors.New("user account is suspended")

	// ErrUserNotFound is returned when a user cannot be found.
	ErrUserNotFound = errors.New("user not found")

	// ErrInvalidStatusTransition is returned when an invalid status transition is attempted.
	ErrInvalidStatusTransition = errors.New("invalid status transition")

	// ErrStatusUnchanged is returned when the status has not changed.
	ErrStatusUnchanged = errors.New("status is already the same")

	ErrUserAlreadyExists = errors.New("user already exists")
)
