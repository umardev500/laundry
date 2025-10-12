package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/order/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"

	paymentMapper "github.com/umardev500/laundry/internal/feature/payment/mapper"
)

// FromEnt converts an Ent Order model to a domain Order.
func FromEnt(e *ent.Order) *domain.Order {
	if e == nil {
		return nil
	}

	order := &domain.Order{
		ID:           e.ID,
		TenantID:     e.TenantID,
		UserID:       e.UserID,
		Status:       types.OrderStatus(e.Status),
		TotalAmount:  e.TotalAmount,
		GuestName:    e.GuestName,
		GuestEmail:   e.GuestEmail,
		GuestPhone:   e.GuestPhone,
		GuestAddress: e.GuestAddress,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
		DeletedAt:    e.DeletedAt,
	}

	// Convert related items if preloaded
	if e.Edges.Items != nil {
		order.Items = FromEntItemList(e.Edges.Items)
	}

	return order
}

// FromEntList converts a list of Ent Orders to domain Orders.
func FromEntList(ents []*ent.Order) []*domain.Order {
	items := make([]*domain.Order, len(ents))
	for i, e := range ents {
		items[i] = FromEnt(e)
	}
	return items
}

// ToResponse converts a domain Order to a response DTO.
func ToResponse(d *domain.Order) *dto.OrderResponse {
	if d == nil {
		return nil
	}

	return &dto.OrderResponse{
		ID:           d.ID,
		TenantID:     d.TenantID,
		UserID:       d.UserID,
		Status:       d.Status,
		TotalAmount:  d.TotalAmount,
		GuestName:    d.GuestName,
		GuestEmail:   d.GuestEmail,
		GuestPhone:   d.GuestPhone,
		GuestAddress: d.GuestAddress,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
		DeletedAt:    d.DeletedAt,
		Items:        ToItemResponseList(d.Items),
		Payment:      paymentMapper.ToResponse(d.Payment),
	}
}

func ToResponseList(orders []*domain.Order) []*dto.OrderResponse {
	res := make([]*dto.OrderResponse, len(orders))
	for i, d := range orders {
		res[i] = ToResponse(d)
	}
	return res
}

// ToResponsePage converts paginated domain.Orders to DTO pagination
func ToResponsePage(data *pagination.PageData[domain.Order]) *pagination.PageData[dto.OrderResponse] {
	if data == nil || len(data.Data) == 0 {
		return &pagination.PageData[dto.OrderResponse]{
			Data:  []*dto.OrderResponse{},
			Total: 0,
		}
	}

	orders := ToResponseList(data.Data)

	return &pagination.PageData[dto.OrderResponse]{
		Data:  orders,
		Total: data.Total,
	}
}
