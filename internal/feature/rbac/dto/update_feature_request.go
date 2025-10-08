package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
)

// UpdateFeatureRequest represents the request body for updating a feature.
type UpdateFeatureRequest struct {
	Name        string `json:"name" validate:"omitempty,min=2,max=100"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

// ToDomain converts the request to a domain.Feature object.
func (r *UpdateFeatureRequest) ToDomain(featureID uuid.UUID) (*domain.Feature, error) {
	feature := &domain.Feature{
		ID:          featureID,
		Name:        r.Name,
		Description: r.Description,
	}
	return feature, nil
}
