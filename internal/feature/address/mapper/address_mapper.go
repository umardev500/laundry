package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/pkg/pagination"

	addressDomain "github.com/umardev500/laundry/internal/feature/address/domain"
	"github.com/umardev500/laundry/internal/feature/address/dto"
	regionDomain "github.com/umardev500/laundry/internal/feature/region/domain"
	regionMapper "github.com/umardev500/laundry/internal/feature/region/mapper"
	userMapper "github.com/umardev500/laundry/internal/feature/user/mapper"
)

// FromEnt converts an Ent Address model to a domain Address.
func FromEnt(e *ent.Addresses) *addressDomain.Address {
	if e == nil {
		return nil
	}

	a := &addressDomain.Address{
		ID:        e.ID,
		User:      userMapper.ToDomainUser(e.Edges.User),
		Street:    e.Street,
		IsPrimary: e.IsPrimary,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}

	// --- Safe region mapping (edges may be nil) ---
	if e.Edges.Province != nil {
		a.Province = &regionDomain.Province{
			ID:   e.Edges.Province.ID,
			Name: e.Edges.Province.Name,
		}
	}

	if e.Edges.Regency != nil {
		a.Regency = &regionDomain.Regency{
			ID:   e.Edges.Regency.ID,
			Name: e.Edges.Regency.Name,
		}
	}

	if e.Edges.District != nil {
		a.District = &regionDomain.District{
			ID:   e.Edges.District.ID,
			Name: e.Edges.District.Name,
		}
	}

	if e.Edges.Village != nil {
		a.Village = &regionDomain.Village{
			ID:   e.Edges.Village.ID,
			Name: e.Edges.Village.Name,
		}
	}

	return a
}

// FromEntList converts a slice of Ent Addresses models to a slice of domain Addresses.
func FromEntList(list []*ent.Addresses) []*addressDomain.Address {
	res := make([]*addressDomain.Address, len(list))
	for i, e := range list {
		res[i] = FromEnt(e)
	}
	return res
}

// ToResponse converts a domain Address to an AddressResponse DTO.
func ToResponse(d *addressDomain.Address) *dto.AddressResponse {
	if d == nil {
		return nil
	}

	return &dto.AddressResponse{
		ID:        d.ID,
		User:      userMapper.ToUserResponse(d.User),
		Province:  regionMapper.ToResponseProvince(d.Province),
		Regency:   regionMapper.ToResponseRegency(d.Regency),
		District:  regionMapper.ToResponseDistrict(d.District),
		Village:   regionMapper.ToResponseVillage(d.Village),
		Street:    d.Street,
		IsPrimary: d.IsPrimary,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		DeletedAt: d.DeletedAt,
	}
}

// ToResponseList converts a slice of domain Addresses to a slice of AddressResponse DTOs.
func ToResponseList(list []*addressDomain.Address) []*dto.AddressResponse {
	res := make([]*dto.AddressResponse, len(list))
	for i, d := range list {
		res[i] = ToResponse(d)
	}
	return res
}

// ToResponsePage converts a paginated domain Address slice to a paginated DTO slice.
func ToResponsePage(data *pagination.PageData[addressDomain.Address]) *pagination.PageData[dto.AddressResponse] {
	res := ToResponseList(data.Data)
	return &pagination.PageData[dto.AddressResponse]{
		Data:  res,
		Total: data.Total,
	}
}
