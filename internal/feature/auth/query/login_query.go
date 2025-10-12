package query

import (
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/internal/app/appctx"
)

type LoginQuery struct {
	Scope appctx.Scope `query:"scope"`
}

func (q *LoginQuery) Validate() error {
	if q.Scope == "" {
		log.Warn().Msg("scope not provided, defaulting to user scope")
		q.Scope = appctx.ScopeUser
		return nil
	}

	if q.Scope != appctx.ScopeUser && q.Scope != appctx.ScopeTenant && q.Scope != appctx.ScopeAdmin {
		return appctx.ErrInvalidScope
	}

	return nil
}
