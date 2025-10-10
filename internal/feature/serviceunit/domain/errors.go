package domain

import "errors"

var (
	ErrServiceUnitNotFound      = errors.New("service unit not found")
	ErrServiceUnitAlreadyExists = errors.New("service unit already exists")
	ErrUnauthorizedAccess       = errors.New("unauthorized access to this service unit")
)
