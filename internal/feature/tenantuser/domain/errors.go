package domain

import "fmt"

var (
	ErrTenantUserAlreadyExists = fmt.Errorf("tenant user already exists")
	ErrTenantOrUserNotFound    = fmt.Errorf("tenant or user not found")
	ErrTenantUserNotFound      = fmt.Errorf("tenant user not found")
	ErrTenantUserDeleted       = fmt.Errorf("tenant user has been deleted")
	ErrInvalidStatusTransition = fmt.Errorf("invalid status transition")
	ErrTenantIDMismatch        = fmt.Errorf("tenant ID mismatch: cannot create tenant user for another tenant")
	ErrUnauthorizedUserAccess  = fmt.Errorf("unauthorized access to user")
)
