package domain

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
)

type Claims struct {
	UserID   uuid.UUID
	TenantID *uuid.UUID
	Scope    appctx.Scope
}
