package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/user/domain"
	"github.com/umardev500/laundry/internal/feature/user/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

// ToDomainUser converts and ent.User to a domain.User
func ToDomainUser(e *ent.User) *domain.User {
	if e == nil {
		return nil
	}

	return &domain.User{
		ID:        e.ID,
		Email:     e.Email,
		Password:  e.Password,
		Status:    domain.Status(*e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}
}

// ToDomainUsers converts a list of ent.User to []*domain.User.
func ToDomainUsers(list []*ent.User) []*domain.User {
	if len(list) == 0 {
		return []*domain.User{}
	}

	users := make([]*domain.User, len(list))
	for i, u := range list {
		users[i] = ToDomainUser(u)
	}
	return users
}

// ToUserResponse converts a domain.User to a dto.UserResponse.
func ToUserResponse(u *domain.User) *dto.UserResponse {
	if u == nil {
		return nil
	}

	return &dto.UserResponse{
		ID:    u.ID,
		Email: u.Email,
	}
}

func ToUserResponsePage(data *pagination.PageData[domain.User]) *pagination.PageData[dto.UserResponse] {
	users := make([]*dto.UserResponse, len(data.Data))
	for i, u := range data.Data {
		users[i] = ToUserResponse(u)
	}

	return &pagination.PageData[dto.UserResponse]{
		Data:  users,
		Total: data.Total,
	}
}
