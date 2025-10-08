package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

type FeatureResponse struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description,omitempty"`
	Status      types.Status `json:"status"`
}
