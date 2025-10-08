package contract

import (
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/auth/domain"
)

type Service interface {
	Login(ctx *appctx.Context, email string, password string) (*domain.LoginResponse, error)
}
