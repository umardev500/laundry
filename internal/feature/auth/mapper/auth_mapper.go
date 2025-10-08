package mapper

import (
	"github.com/umardev500/laundry/internal/feature/auth/domain"
	"github.com/umardev500/laundry/internal/feature/auth/dto"
)

func FromDomain(d *domain.LoginResponse) *dto.LoginResponse {
	if d == nil {
		return nil
	}

	return &dto.LoginResponse{
		Tokens: dto.Tokens{
			AccessToken:  d.Tokens.AccessToken,
			RefreshToken: d.Tokens.RefreshToken,
			ExpiresAt:    d.Tokens.ExpiresAt,
		},
	}
}
