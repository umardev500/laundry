package appctx

type ContextKey string

const (
	ContextKeyTenantID ContextKey = "tenant_id"
	ContextKeyUserID   ContextKey = "user_id"
	ContextKeyScope    ContextKey = "scope"
)
