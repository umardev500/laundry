package appctx

type Scope string

const (
	ScopeUser   Scope = "user"
	ScopeAdmin  Scope = "admin"
	ScopeTenant Scope = "tenant"
)
