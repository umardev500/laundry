package dto

import (
	"github.com/google/uuid"
)

type TenantResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone,omitempty"`
	Address string    `json:"address,omitempty"`
	Status  string    `json:"status"`
}
