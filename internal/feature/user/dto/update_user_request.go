package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/user/domain"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserRequest struct {
	Email    *string `json:"email" validate:"omitempty,email"`
	Password *string `json:"password"`
}

func (r *UpdateUserRequest) HashPassword() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(*r.Password), bcrypt.DefaultCost)
}

func (r *UpdateUserRequest) ToDomainUserWithID(uid uuid.UUID) (*domain.User, error) {
	hashed, err := r.HashPassword()
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       uid,
		Email:    *r.Email,
		Password: string(hashed),
		Status:   domain.StatusActive,
	}, nil
}
