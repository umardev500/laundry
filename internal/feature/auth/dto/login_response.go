package dto

import "time"

type Tokens struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type LoginResponse struct {
	Tokens Tokens `json:"tokens"`
}
