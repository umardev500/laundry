package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/address/domain"
	regionDomain "github.com/umardev500/laundry/internal/feature/region/domain"
)

type UpdateAddressRequest struct {
	ProvinceID *string `json:"province_id,omitempty"`
	RegencyID  *string `json:"regency_id,omitempty"`
	DistrictID *string `json:"district_id,omitempty"`
	VillageID  *string `json:"village_id,omitempty"`
	Street     *string `json:"street,omitempty"`
	IsPrimary  *bool   `json:"is_primary,omitempty"`
}

func (r *UpdateAddressRequest) ToDomain(id uuid.UUID) *domain.Address {
	in := &domain.Address{
		ID: id,
	}

	if r.ProvinceID != nil && *r.ProvinceID != "" {
		in.Province = &regionDomain.Province{ID: *r.ProvinceID}
	}

	if r.RegencyID != nil && *r.RegencyID != "" {
		in.Regency = &regionDomain.Regency{ID: *r.RegencyID}
	}

	if r.DistrictID != nil && *r.DistrictID != "" {
		in.District = &regionDomain.District{ID: *r.DistrictID}
	}

	if r.VillageID != nil && *r.VillageID != "" {
		in.Village = &regionDomain.Village{ID: *r.VillageID}
	}

	if r.Street != nil {
		in.Street = r.Street
	}

	if r.IsPrimary != nil {
		in.IsPrimary = *r.IsPrimary
	}

	return in
}
