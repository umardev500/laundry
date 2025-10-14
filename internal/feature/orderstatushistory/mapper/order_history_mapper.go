package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/domain"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

// FromEnt converts Ent OrderStatusHistory to domain struct
func FromEntStatusHistory(e *ent.OrderStatusHistory) *domain.OrderStatusHistory {
	if e == nil {
		return nil
	}

	ref := getRef(e)

	return &domain.OrderStatusHistory{
		ID:        e.ID,
		OrderID:   e.OrderID,
		Status:    types.OrderStatus(e.Status), // or types.OrderStatus depending on your domain
		Notes:     e.Notes,
		CreatedAt: e.CreatedAt,
		Order:     ref,
	}
}

// FromEntStatusHistoryList converts a list of Ent OrderStatusHistory to domain structs
func FromEntStatusHistoryList(list []*ent.OrderStatusHistory) []*domain.OrderStatusHistory {
	res := make([]*domain.OrderStatusHistory, len(list))
	for i, e := range list {
		res[i] = FromEntStatusHistory(e)
	}
	return res
}

// ToResponse converts domain OrderStatusHistory to DTO for API response
func ToResponse(d *domain.OrderStatusHistory, refMapper types.RefMapper) *dto.OrderStatusHistoryResponse {
	if d == nil {
		return nil
	}
	return &dto.OrderStatusHistoryResponse{
		ID:        d.ID,
		OrderID:   d.OrderID,
		Status:    d.Status,
		Notes:     d.Notes,
		CreatedAt: d.CreatedAt,
		Order:     refMapper(d.Order),
	}
}

// ToResponsePage converts paginated domain OrderStatusHistory to DTO page
func ToResponsePage(data *pagination.PageData[domain.OrderStatusHistory], refMapper types.RefMapper) *pagination.PageData[dto.OrderStatusHistoryResponse] {
	res := make([]*dto.OrderStatusHistoryResponse, len(data.Data))
	for i, m := range data.Data {
		res[i] = ToResponse(m, refMapper)
	}
	return &pagination.PageData[dto.OrderStatusHistoryResponse]{
		Data:  res,
		Total: data.Total,
	}
}

func getRef(e *ent.OrderStatusHistory) any {
	if e == nil {
		return nil
	}
	return e.Edges.Order
}
