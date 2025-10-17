package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/subscription/domain"
	"github.com/umardev500/laundry/internal/feature/subscription/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

// FromEnt converts an Ent Subscription model to a domain Subscription.
func FromEnt(e *ent.Subscription) *domain.Subscription {
	if e == nil {
		return nil
	}

	return &domain.Subscription{
		ID:        e.ID,
		TenantID:  e.TenantID,
		PlanID:    e.PlanID,
		Status:    types.SubscriptionStatus(e.Status),
		StartDate: e.StartDate,
		EndDate:   e.EndDate,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}
}

// FromEntList converts a slice of Ent Subscription models to a slice of domain Subscriptions.
func FromEntList(list []*ent.Subscription) []*domain.Subscription {
	res := make([]*domain.Subscription, len(list))
	for i, e := range list {
		res[i] = FromEnt(e)
	}
	return res
}

// ToResponse converts a domain Subscription to a SubscriptionResponse DTO.
func ToResponse(d *domain.Subscription) *dto.SubscriptionResponse {
	if d == nil {
		return nil
	}

	return &dto.SubscriptionResponse{
		ID:        d.ID,
		TenantID:  d.TenantID,
		PlanID:    d.PlanID,
		Status:    d.Status,
		StartDate: d.StartDate,
		EndDate:   d.EndDate,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		DeletedAt: d.DeletedAt,
	}
}

// ToResponseList converts a slice of domain Subscriptions to a slice of SubscriptionResponse DTOs.
func ToResponseList(list []*domain.Subscription) []*dto.SubscriptionResponse {
	res := make([]*dto.SubscriptionResponse, len(list))
	for i, d := range list {
		res[i] = ToResponse(d)
	}
	return res
}

// ToResponsePage converts a paginated domain Subscription slice to a paginated DTO slice.
func ToResponsePage(data *pagination.PageData[domain.Subscription]) *pagination.PageData[dto.SubscriptionResponse] {
	res := ToResponseList(data.Data)
	return &pagination.PageData[dto.SubscriptionResponse]{
		Data:  res,
		Total: data.Total,
	}
}
