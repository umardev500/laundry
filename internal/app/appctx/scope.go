package appctx

import "fmt"

type Scope string

const (
	ScopeUser   Scope = "user"
	ScopeAdmin  Scope = "admin"
	ScopeTenant Scope = "tenant"
)

var (
	ErrInvalidScope = fmt.Errorf("invalid scope")
)
