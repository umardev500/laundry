package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

type MachineResponse struct {
	ID            uuid.UUID           `json:"id"`
	TenantID      uuid.UUID           `json:"tenant_id"`
	MachineTypeID uuid.UUID           `json:"machine_type_id"`
	Name          string              `json:"name"`
	Description   string              `json:"description,omitempty"`
	Status        types.MachineStatus `json:"status"`
}
