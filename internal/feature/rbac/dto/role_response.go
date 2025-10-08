package dto

import "github.com/google/uuid"

type RoleResponse struct {
	ID          string     `json:"id"`
	TenantID    *uuid.UUID `json:"tenant_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
}
