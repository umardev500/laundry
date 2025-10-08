package domain

import (
	"fmt"
)

var (
	ErrRoleAlreadyExists      = fmt.Errorf("role already exists")
	ErrRoleNotFound           = fmt.Errorf("role not found")
	ErrRoleDeleted            = fmt.Errorf("role has been deleted")
	ErrEmptyRoleName          = fmt.Errorf("role name cannot be empty")
	ErrUnauthorizedRoleAccess = fmt.Errorf("unauthorized access to role")
	ErrMissingTenantID        = fmt.Errorf("tenant ID cannot be nil")
)
