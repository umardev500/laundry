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

var (
	ErrFeatureNotFound  = fmt.Errorf("feature not found")
	ErrFeatureDeleted   = fmt.Errorf("feature has been deleted")
	ErrFeatureImmutable = fmt.Errorf("feature cannot be created or deleted manually")
)

var (
	ErrPermissionAlreadyExists = fmt.Errorf("permission already exists")
	ErrPermissionNotFound      = fmt.Errorf("permission not found")
	ErrPermissionDeleted       = fmt.Errorf("permission has been deleted")
)
