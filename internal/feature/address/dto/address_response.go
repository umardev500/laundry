package dto

import (
	"time"

	"github.com/google/uuid"

	regionDto "github.com/umardev500/laundry/internal/feature/region/dto"
	userDto "github.com/umardev500/laundry/internal/feature/user/dto"
)

type AddressResponse struct {
	ID        uuid.UUID                   `json:"id"`
	User      *userDto.UserResponse       `json:"user,omitempty"`
	Province  *regionDto.ProvinceResponse `json:"province,omitempty"`
	Regency   *regionDto.RegencyResponse  `json:"regency,omitempty"`
	District  *regionDto.DistrictResponse `json:"district,omitempty"`
	Village   *regionDto.VillageResponse  `json:"village,omitempty"`
	Street    *string                     `json:"street,omitempty"`
	IsPrimary bool                        `json:"is_primary"`
	CreatedAt time.Time                   `json:"created_at"`
	UpdatedAt time.Time                   `json:"updated_at"`
	DeletedAt *time.Time                  `json:"deleted_at,omitempty"`
}
