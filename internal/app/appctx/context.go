package appctx

import (
	"context"

	"github.com/google/uuid"
)

type Context struct {
	context.Context
}

// --- Constructors ---
func New(ctx context.Context) *Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Context{
		Context: ctx,
	}
}

// --- Tenant ---
func (c *Context) TenantID() *uuid.UUID {
	val := c.Value(ContextKeyTenantID)
	if val != nil {
		return val.(*uuid.UUID)
	}

	return nil
}

func (c *Context) WithTenantID(id *uuid.UUID) *Context {
	return &Context{context.WithValue(c.Context, ContextKeyTenantID, id)}
}

// --- User ---
func (c *Context) UserID() *uuid.UUID {
	val := c.Value(ContextKeyUserID)
	if val != nil {
		if id, ok := val.(*uuid.UUID); ok {
			return id
		}
	}
	return nil
}

func (c *Context) WithUserID(id *uuid.UUID) *Context {
	return &Context{context.WithValue(c.Context, ContextKeyUserID, id)}
}

// --- Scope ---
func (c *Context) Scope() Scope {
	val := c.Value(ContextKeyScope)
	if val != nil {
		if s, ok := val.(Scope); ok {
			return s
		}
	}
	return ""
}

func (c *Context) WithScope(s Scope) *Context {
	return &Context{context.WithValue(c.Context, ContextKeyScope, s)}
}
