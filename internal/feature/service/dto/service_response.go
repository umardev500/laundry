package dto

import (
	"time"

	"github.com/google/uuid"
)

type ServiceResponse struct {
	ID                uuid.UUID  `json:"id"`
	TenantID          uuid.UUID  `json:"tenant_id"`
	ServiceUnitID     *uuid.UUID `json:"service_unit_id,omitempty"`
	ServiceCategoryID *uuid.UUID `json:"service_category_id,omitempty"`
	Name              string     `json:"name"`
	Price             float64    `json:"price"`
	Description       string     `json:"description,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
