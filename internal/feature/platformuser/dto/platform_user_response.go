package dto

import "github.com/google/uuid"

// PlatformUserResponse represents the API response for a PlatformUser
type PlatformUserResponse struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Status string    `json:"status"`
}
