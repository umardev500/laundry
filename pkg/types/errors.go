package types

import "errors"

var (
	ErrStatusUnchanged         = errors.New("status is already the same")
	ErrInvalidStatusTransition = errors.New("invalid status transition")
	ErrInvalidStatus           = errors.New("invalid status value: must be 'active' or 'suspended'")
	ErrTenantIDRequired        = errors.New("tenant ID is required")
	ErrInvalidUUID             = errors.New("invalid UUID")
)
