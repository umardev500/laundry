package dto

import (
	"errors"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/platformuser/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// CreatePlatformUserRequest represents the payload to create a PlatformUser
type CreatePlatformUserRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

// ToDomain converts CreatePlatformUserRequest to domain.PlatformUser
func (r *CreatePlatformUserRequest) ToDomain() (*domain.PlatformUser, error) {
	if r.UserID == uuid.Nil {
		return nil, errors.New("user_id is required")
	}

	return &domain.PlatformUser{
		UserID: r.UserID,
		Status: types.StatusActive, // default active
	}, nil
}
