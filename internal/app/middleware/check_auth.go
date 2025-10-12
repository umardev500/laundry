package middleware

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/auth/domain"
	"github.com/umardev500/laundry/pkg/httpx"
)

func CheckAuth(config *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return httpx.Unauthorized(c, "missing authorization header")
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return httpx.Unauthorized(c, "invalid authorization header format")
		}

		token, err := jwt.Parse([]byte(parts[1]), jwt.WithKey(jwa.HS256(), []byte(config.JWT.Secret)))
		if err != nil {
			return httpx.Unauthorized(c, "invalid token")
		}

		// Extract claims
		sub, ok := token.Subject()
		if !ok {
			return httpx.Unauthorized(c, "invalid token")
		}

		var scope string
		if err := token.Get(string(appctx.ContextKeyScope), &scope); err != nil {
			return httpx.Unauthorized(c, "invalid token")
		}

		var tenantIDStr string
		var tenantID *uuid.UUID

		if err := token.Get(string(appctx.ContextKeyTenantID), &tenantIDStr); err == nil {
			if id, err := uuid.Parse(tenantIDStr); err == nil {
				tenantID = &id
			}
		}

		claims := &domain.Claims{
			UserID:   uuid.MustParse(sub),
			Scope:    appctx.Scope(scope),
			TenantID: tenantID,
		}

		ctx := context.Background()
		ctx = context.WithValue(ctx, appctx.ContextKeyUserID, &claims.UserID)
		ctx = context.WithValue(ctx, appctx.ContextKeyScope, claims.Scope)
		ctx = context.WithValue(ctx, appctx.ContextKeyTenantID, claims.TenantID)
		c.SetUserContext(ctx)

		return c.Next()
	}
}
