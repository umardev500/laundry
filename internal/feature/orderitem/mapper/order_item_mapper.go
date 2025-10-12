package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/orderitem/domain"
	"github.com/umardev500/laundry/internal/feature/orderitem/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

// FromEnt converts an Ent OrderItem model to a domain OrderItem.
func FromEnt(e *ent.OrderItem) *domain.OrderItem {
	if e == nil {
		return nil
	}

	return &domain.OrderItem{
		ID:          e.ID,
		OrderID:     e.OrderID,
		ServiceID:   e.ServiceID,
		Quantity:    e.Quantity,
		Price:       e.Price,
		Subtotal:    e.Subtotal,
		TotalAmount: e.TotalAmount,
	}
}

// FromEntList converts a list of Ent OrderItems to domain OrderItems.
func FromEntList(ents []*ent.OrderItem) []*domain.OrderItem {
	items := make([]*domain.OrderItem, len(ents))
	for i, e := range ents {
		items[i] = FromEnt(e)
	}
	return items
}

// ToResponse converts a domain OrderItem to a response DTO.
func ToResponse(d *domain.OrderItem) *dto.OrderItemResponse {
	if d == nil {
		return nil
	}

	return &dto.OrderItemResponse{
		ID:          d.ID,
		OrderID:     d.OrderID,
		ServiceID:   d.ServiceID,
		Quantity:    d.Quantity,
		Price:       d.Price,
		Subtotal:    d.Subtotal,
		TotalAmount: d.TotalAmount,
	}
}

// ToResponseList converts a list of domain OrderItems to response DTOs.
func ToResponseList(items []*domain.OrderItem) []*dto.OrderItemResponse {
	res := make([]*dto.OrderItemResponse, len(items))
	for i, item := range items {
		res[i] = ToResponse(item)
	}
	return res
}

// ToResponsePage converts paginated domain.OrderItems to DTO pagination.
func ToResponsePage(data *pagination.PageData[domain.OrderItem]) *pagination.PageData[dto.OrderItemResponse] {
	if data == nil || len(data.Data) == 0 {
		return &pagination.PageData[dto.OrderItemResponse]{
			Data:  []*dto.OrderItemResponse{},
			Total: 0,
		}
	}

	items := ToResponseList(data.Data)

	return &pagination.PageData[dto.OrderItemResponse]{
		Data:  items,
		Total: data.Total,
	}
}
