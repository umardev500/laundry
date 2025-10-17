package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// SubscriptionResponse represents a subscription returned via API.
type SubscriptionResponse struct {
	ID        uuid.UUID                `json:"id"`
	TenantID  uuid.UUID                `json:"tenant_id"`
	PlanID    uuid.UUID                `json:"plan_id"`
	Status    types.SubscriptionStatus `json:"status"`
	StartDate *time.Time               `json:"start_date,omitempty"`
	EndDate   *time.Time               `json:"end_date,omitempty"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
	DeletedAt *time.Time               `json:"deleted_at,omitempty"`
}
