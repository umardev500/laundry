package domain

import "time"

// Tokens holds both access and refresh tokens
type Tokens struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

// LoginResponse represents the result of a successful login
type LoginResponse struct {
	Tokens Tokens
}
