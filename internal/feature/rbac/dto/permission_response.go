package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// PermissionResponse represents how a permission is returned to the client.
type PermissionResponse struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	DisplayName string       `json:"display_name"`
	Description string       `json:"description,omitempty"`
	Status      types.Status `json:"status"`
	FeatureID   uuid.UUID    `json:"feature_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   *time.Time   `json:"deleted_at,omitempty"`
}
