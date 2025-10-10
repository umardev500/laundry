package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// Machine holds the schema definition for the Machine entity.
type Machine struct {
	ent.Schema
}

// Fields of the Machine.
func (Machine) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("tenant_id", uuid.UUID{}).Immutable(),
		field.UUID("machine_type_id", uuid.UUID{}).Optional().Nillable(),
		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional(),
		field.Enum("status").
			Values(
				string(types.MachineStatusAvailable),
				string(types.MachineStatusInUse),
				string(types.MachineStatusMaintenance),
				string(types.MachineStatusOffline),
				string(types.MachineStatusReserved),
			).
			Default(string(types.MachineStatusAvailable)),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Machine.
func (Machine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("machine_type", MachineType.Type).
			Ref("machines").
			Field("machine_type_id").
			Unique(),
	}
}
