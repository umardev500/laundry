package domain

import "errors"

var (
	ErrPlatformUserAlreadyExists = errors.New("platform user already exists")
	ErrPlatformUserNotFound      = errors.New("platform user not found")
	ErrPlatformUserDeleted       = errors.New("platform user deleted")
	ErrStatusUnchanged           = errors.New("status is already the same")
	ErrInvalidStatusTransition   = errors.New("invalid status transition")
)
