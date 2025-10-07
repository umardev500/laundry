package appctx

import "context"

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
