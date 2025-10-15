package domain

import "fmt"

var (
	ErrTenantAlreadyExists = fmt.Errorf("tenant already exists")
	ErrTenantNotFound      = fmt.Errorf("tenant not found")
	ErrTenantDeleted       = fmt.Errorf("tenant has been deleted")
)
