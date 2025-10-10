package dto

import (
	"github.com/google/uuid"
)

type MachineTypeResponse struct {
	ID          uuid.UUID `json:"id"`
	TenantID    uuid.UUID `json:"tenant_id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Capacity    int       `json:"capacity,omitempty"`
}
