package domain

import "fmt"

var (
	ErrTenantUserAlreadyExists = fmt.Errorf("tenant user already exists")
	ErrTenantUserNotFound      = fmt.Errorf("tenant user not found")
	ErrTenantUserDeleted       = fmt.Errorf("tenant user has been deleted")
	ErrStatusUnchanged         = fmt.Errorf("status unchanged")
	ErrInvalidStatusTransition = fmt.Errorf("invalid status transition")
)
