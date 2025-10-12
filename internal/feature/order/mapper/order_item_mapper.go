package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/order/dto"
	"github.com/umardev500/laundry/pkg/pagination"

	orderItemDomain "github.com/umardev500/laundry/internal/feature/orderitem/domain"
)

// FromEntItem converts an Ent OrderItem to a domain OrderItem.
func FromEntItem(e *ent.OrderItem) *orderItemDomain.OrderItem {
	if e == nil {
		return nil
	}

	return &orderItemDomain.OrderItem{
		ID:          e.ID,
		OrderID:     e.OrderID,
		ServiceID:   e.ServiceID,
		Quantity:    e.Quantity,
		Price:       e.Price,
		Subtotal:    e.Subtotal,
		TotalAmount: e.TotalAmount,
	}
}

// FromEntItemList converts a list of Ent OrderItems to domain OrderItems.
func FromEntItemList(ents []*ent.OrderItem) []*orderItemDomain.OrderItem {
	items := make([]*orderItemDomain.OrderItem, len(ents))
	for i, e := range ents {
		items[i] = FromEntItem(e)
	}
	return items
}

// ToItemResponse converts a domain OrderItem to a response DTO.
func ToItemResponse(d *orderItemDomain.OrderItem) *dto.OrderItemResponse {
	if d == nil {
		return nil
	}

	return &dto.OrderItemResponse{
		ID:        d.ID,
		OrderID:   d.OrderID,
		ServiceID: d.ServiceID,
		Quantity:  d.Quantity,
		Price:     d.Price,
		Total:     d.TotalAmount,
	}
}

// ToItemResponseList converts a list of domain OrderItems to response DTOs.
func ToItemResponseList(items []*orderItemDomain.OrderItem) []*dto.OrderItemResponse {
	res := make([]*dto.OrderItemResponse, len(items))
	for i, d := range items {
		res[i] = ToItemResponse(d)
	}
	return res
}

// ToItemResponsePage converts paginated orderItemDomain.OrderItems to paginated DTOs.
func ToItemResponsePage(data *pagination.PageData[orderItemDomain.OrderItem]) *pagination.PageData[dto.OrderItemResponse] {
	if data == nil || len(data.Data) == 0 {
		return &pagination.PageData[dto.OrderItemResponse]{
			Data:  []*dto.OrderItemResponse{},
			Total: 0,
		}
	}

	items := ToItemResponseList(data.Data)

	return &pagination.PageData[dto.OrderItemResponse]{
		Data:  items,
		Total: data.Total,
	}
}
