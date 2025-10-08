package domain

type Scope string

const (
	ScopeUser  Scope = "user"
	ScopeAdmin Scope = "admin"
)

type ClaimKey string

const (
	ClaimScope ClaimKey = "scope"
)
