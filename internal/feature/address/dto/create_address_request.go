package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/address/domain"
)

type CreateAddressRequest struct {
	ProvinceID string  `json:"province_id" validate:"required"`
	RegencyID  string  `json:"regency_id" validate:"required"`
	DistrictID string  `json:"district_id" validate:"required"`
	VillageID  string  `json:"village_id" validate:"required"`
	Street     *string `json:"street,omitempty"`
	IsPrimary  *bool   `json:"is_primary"`
}

func (r *CreateAddressRequest) ToDomain(userID uuid.UUID) (*domain.Address, error) {
	return domain.NewAddress(
		userID,
		r.ProvinceID,
		r.RegencyID,
		r.DistrictID,
		r.VillageID,
		r.Street,
		r.IsPrimary,
	)
}
