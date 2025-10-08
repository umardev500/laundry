package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/platformuser"
	"github.com/umardev500/laundry/internal/feature/platformuser/domain"
	"github.com/umardev500/laundry/internal/feature/platformuser/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

func FromEntModel(e *ent.PlatformUser) *domain.PlatformUser {
	return &domain.PlatformUser{
		ID:        e.ID,
		UserID:    e.UserID,
		Status:    types.Status(*e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}
}

func FromEntModels(es []*ent.PlatformUser) []*domain.PlatformUser {
	var res []*domain.PlatformUser
	for _, e := range es {
		res = append(res, FromEntModel(e))
	}
	return res
}

func ToEntModel(d *domain.PlatformUser) *ent.PlatformUser {
	return &ent.PlatformUser{
		ID:        d.ID,
		UserID:    d.UserID,
		Status:    (*platformuser.Status)(&d.Status),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		DeletedAt: d.DeletedAt,
	}
}

// ToPlatformUserResponse converts domain.PlatformUser to dto.PlatformUserResponse
func ToPlatformUserResponse(d *domain.PlatformUser) *dto.PlatformUserResponse {
	if d == nil {
		return nil
	}
	return &dto.PlatformUserResponse{
		ID:     d.ID,
		UserID: d.UserID,
		Status: string(d.Status),
	}
}

// ToPlatformUserResponsePage converts pagination.PageData[domain.PlatformUser] to pagination.PageData[dto.PlatformUserResponse]
func ToPlatformUserResponsePage(data *pagination.PageData[domain.PlatformUser]) *pagination.PageData[dto.PlatformUserResponse] {
	res := make([]*dto.PlatformUserResponse, len(data.Data))
	for i, pu := range data.Data {
		res[i] = ToPlatformUserResponse(pu)
	}

	return &pagination.PageData[dto.PlatformUserResponse]{
		Data:  res,
		Total: data.Total,
	}
}
