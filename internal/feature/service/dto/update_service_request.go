package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/service/domain"
	"github.com/umardev500/laundry/pkg/utils"
)

type UpdateServiceRequest struct {
	Name              string    `json:"name,omitempty" validate:"omitempty,min=2,max=200"`
	Price             *float64  `json:"price,omitempty"`
	Description       string    `json:"description,omitempty" validate:"omitempty,max=1024"`
	ServiceUnitID     uuid.UUID `json:"service_unit_id,omitempty"`
	ServiceCategoryID uuid.UUID `json:"service_category_id,omitempty"`
}

func (r *UpdateServiceRequest) ToDomain(id uuid.UUID) *domain.Service {
	price := float64(-1)
	if r.Price != nil {
		price = *r.Price
	}
	return &domain.Service{
		ID:                id,
		Name:              r.Name,
		BasePrice:         price,
		Description:       r.Description,
		ServiceUnitID:     utils.NilIfUUIDZero(r.ServiceUnitID),
		ServiceCategoryID: utils.NilIfUUIDZero(r.ServiceCategoryID),
	}
}
