package service

import (
	"fmt"
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

	tenantUserContract "github.com/umardev500/laundry/internal/feature/tenantuser/contract"
	userContract "github.com/umardev500/laundry/internal/feature/user/contract"
)

type serviceImpl struct {
	config            *config.Config
	userService       userContract.Service
	refreshTokenRepo  repository.RefreshTokenRepository
	tenantUserService tenantUserContract.Service
}

func NewService(
	cfg *config.Config,
	userService userContract.Service,
	refreshTokenRepo repository.RefreshTokenRepository,
	tenantService tenantUserContract.Service,
) contract.Service {
	return &serviceImpl{
		config:            cfg,
		userService:       userService,
		refreshTokenRepo:  refreshTokenRepo,
		tenantUserService: tenantService,
	}
}

// LoginTenant implements contract.Service.
func (s *serviceImpl) LoginTenant(ctx *appctx.Context, email string, password string) (*domain.LoginResponse, error) {
	email = normalizeEmail(email)

	user, err := s.userService.GetByEmail(ctx, email)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	if !security.Compare(user.Password, password) {
		return nil, domain.ErrInvalidCredentials
	}

	// Get tenant
	tenants, err := s.tenantUserService.GetByUser(ctx, user.ID)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	if len(tenants) == 0 {
		return nil, domain.ErrInvalidCredentials
	}

	if len(tenants) > 1 {
		return nil, domain.ErrMultipleTenants
	}

	tenant := tenants[0]

	// Build JWT claims
	now := time.Now().UTC()
	jwtCfg := s.config.JWT
	exp := now.Add(time.Duration(jwtCfg.ExpirySeconds) * time.Second)

	tokenBuilder := jwt.NewBuilder().
		Issuer(jwtCfg.Issuer).
		Subject(user.ID.String()).
		IssuedAt(now).
		Expiration(exp).
		Claim(string(appctx.ContextKeyScope), appctx.ScopeTenant).
		Claim(string(appctx.ContextKeyTenantID), tenant.TenantID.String())

	token, err := tokenBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf("build jwt token: %w", err)
	}

	// Sign JWT
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256(), []byte(jwtCfg.Secret)))
	if err != nil {
		return nil, fmt.Errorf("sign jwt token: %w", err)
	}

	// Generate and store refresh token
	refreshToken, err := s.issueRefreshToken(ctx, user.ID.String())
	if err != nil {
		return nil, fmt.Errorf("issue refresh token: %w", err)
	}

	return &domain.LoginResponse{
		Tokens: domain.Tokens{
			AccessToken:  string(signed),
			RefreshToken: refreshToken,
			ExpiresAt:    exp,
		},
	}, nil
}

// Login authenticates a user and issues access & refresh tokens.
func (s *serviceImpl) Login(ctx *appctx.Context, email, password string) (*domain.LoginResponse, error) {
	email = normalizeEmail(email)

	user, err := s.userService.GetByEmail(ctx, email)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	if !security.Compare(user.Password, password) {
		return nil, domain.ErrInvalidCredentials
	}

	// Build JWT claims
	now := time.Now().UTC()
	jwtCfg := s.config.JWT
	exp := now.Add(time.Duration(jwtCfg.ExpirySeconds) * time.Second)

	tokenBuilder := jwt.NewBuilder().
		Issuer(jwtCfg.Issuer).
		Subject(user.ID.String()).
		IssuedAt(now).
		Expiration(exp).
		Claim(string(appctx.ContextKeyScope), appctx.ScopeUser)

	token, err := tokenBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf("build jwt token: %w", err)
	}

	// Sign JWT
	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256(), []byte(jwtCfg.Secret)))
	if err != nil {
		return nil, fmt.Errorf("sign jwt token: %w", err)
	}

	// Generate and store refresh token
	refreshToken, err := s.issueRefreshToken(ctx, user.ID.String())
	if err != nil {
		return nil, fmt.Errorf("issue refresh token: %w", err)
	}

	return &domain.LoginResponse{
		Tokens: domain.Tokens{
			AccessToken:  string(signed),
			RefreshToken: refreshToken,
			ExpiresAt:    exp,
		},
	}, nil
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
