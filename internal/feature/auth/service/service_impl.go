package service

import (
	"fmt"
	"maps"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/auth/contract"
	"github.com/umardev500/laundry/internal/feature/auth/domain"
	"github.com/umardev500/laundry/internal/feature/auth/repository"
	"github.com/umardev500/laundry/pkg/security"

	platformUserContract "github.com/umardev500/laundry/internal/feature/platformuser/contract"
	tenantUserContract "github.com/umardev500/laundry/internal/feature/tenantuser/contract"
	userContract "github.com/umardev500/laundry/internal/feature/user/contract"
)

type serviceImpl struct {
	config              *config.Config
	userService         userContract.Service
	refreshTokenRepo    repository.RefreshTokenRepository
	tenantUserService   tenantUserContract.Service
	platformUserService platformUserContract.Service
}

func NewService(
	cfg *config.Config,
	userService userContract.Service,
	refreshTokenRepo repository.RefreshTokenRepository,
	tenantService tenantUserContract.Service,
	platformUserService platformUserContract.Service,
) contract.Service {
	return &serviceImpl{
		config:              cfg,
		userService:         userService,
		refreshTokenRepo:    refreshTokenRepo,
		tenantUserService:   tenantService,
		platformUserService: platformUserService,
	}
}

// --- Public login methods ---

func (s *serviceImpl) LoginAdmin(ctx *appctx.Context, email, password string) (*domain.LoginResponse, error) {
	return s.loginWithScope(ctx, email, password, appctx.ScopeAdmin, func(userID uuid.UUID) (map[string]any, error) {
		pu, err := s.platformUserService.GetByUserID(ctx, userID)
		if err != nil || pu == nil {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, nil // No extra claims needed for admin
	})
}

func (s *serviceImpl) LoginTenant(ctx *appctx.Context, email, password string) (*domain.LoginResponse, error) {
	return s.loginWithScope(ctx, email, password, appctx.ScopeTenant, func(userID uuid.UUID) (map[string]any, error) {
		tenants, err := s.tenantUserService.GetByUser(ctx, userID)
		if err != nil {
			return nil, domain.ErrInvalidCredentials
		}
		if len(tenants) == 0 {
			return nil, domain.ErrInvalidCredentials
		}
		if len(tenants) > 1 {
			return nil, domain.ErrMultipleTenants
		}
		return map[string]any{
			string(appctx.ContextKeyTenantID): tenants[0].TenantID.String(),
		}, nil
	})
}

func (s *serviceImpl) Login(ctx *appctx.Context, email, password string) (*domain.LoginResponse, error) {
	return s.loginWithScope(ctx, email, password, appctx.ScopeUser, nil)
}

// --- Internal helper ---

type extraClaimsFunc func(userID uuid.UUID) (map[string]any, error)

func (s *serviceImpl) loginWithScope(ctx *appctx.Context, email, password string, scope appctx.Scope, extraClaims extraClaimsFunc) (*domain.LoginResponse, error) {
	email = normalizeEmail(email)

	user, err := s.userService.GetByEmail(ctx, email)
	if err != nil || !security.Compare(user.Password, password) {
		return nil, domain.ErrInvalidCredentials
	}

	claims := map[string]any{
		string(appctx.ContextKeyScope): scope,
	}

	if extraClaims != nil {
		extra, err := extraClaims(user.ID)
		if err != nil {
			return nil, err
		}

		maps.Copy(claims, extra)
	}

	accessToken, exp, err := s.buildJWT(user.ID.String(), claims)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.issueRefreshToken(ctx, user.ID.String())
	if err != nil {
		return nil, fmt.Errorf("issue refresh token: %w", err)
	}

	return &domain.LoginResponse{
		Tokens: domain.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresAt:    exp,
		},
	}, nil
}

// buildJWT creates and signs a JWT token.
func (s *serviceImpl) buildJWT(userID string, claims map[string]any) (string, time.Time, error) {
	now := time.Now().UTC()
	exp := now.Add(time.Duration(s.config.JWT.ExpirySeconds) * time.Second)

	builder := jwt.NewBuilder().
		Issuer(s.config.JWT.Issuer).
		Subject(userID).
		IssuedAt(now).
		Expiration(exp)

	for k, v := range claims {
		builder.Claim(k, v)
	}

	token, err := builder.Build()
	if err != nil {
		return "", time.Time{}, fmt.Errorf("build jwt token: %w", err)
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256(), []byte(s.config.JWT.Secret)))
	if err != nil {
		return "", time.Time{}, fmt.Errorf("sign jwt token: %w", err)
	}

	return string(signed), exp, nil
}

// issueRefreshToken creates and persists a refresh token securely.
func (s *serviceImpl) issueRefreshToken(ctx *appctx.Context, userID string) (string, error) {
	rt := uuid.NewString()
	now := time.Now().UTC()
	exp := now.Add(time.Duration(s.config.JWT.RefreshTokenExpirySeconds) * time.Second)

	if err := s.refreshTokenRepo.Set(ctx, userID, rt, exp.Sub(now)); err != nil {
		return "", err
	}
	return rt, nil
}

// normalizeEmail lowercases and trims spaces.
func normalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
