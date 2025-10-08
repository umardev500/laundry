package types

import "errors"

var (
	ErrStatusUnchanged         = errors.New("status is already the same")
	ErrInvalidStatusTransition = errors.New("invalid status transition")
	ErrInvalidStatus           = errors.New("invalid status value: must be 'active' or 'suspended'")
)
