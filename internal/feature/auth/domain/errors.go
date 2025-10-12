package domain

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrMultipleTenants    = errors.New("multiple tenants found for user")
)
