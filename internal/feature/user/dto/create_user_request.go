package dto

import (
	"github.com/umardev500/laundry/internal/feature/user/domain"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"password"`
}

func (r *CreateUserRequest) HashPassword() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
}

func (r *CreateUserRequest) ToDomainUser() (*domain.User, error) {
	hashed, err := r.HashPassword()
	if err != nil {
		return nil, err
	}

	return &domain.User{
		Email:    r.Email,
		Password: string(hashed),
		Status:   domain.StatusActive,
	}, nil
}
