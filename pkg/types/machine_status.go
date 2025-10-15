package types

import (
	"slices"
	"strings"
)

// MachineStatus represents the lifecycle state of a machine.
type MachineStatus string

const (
	MachineStatusAvailable   MachineStatus = "AVAILABLE"
	MachineStatusInUse       MachineStatus = "IN_USE"
	MachineStatusMaintenance MachineStatus = "MAINTENANCE"
	MachineStatusOffline     MachineStatus = "OFFLINE"
	MachineStatusReserved    MachineStatus = "RESERVED"
)

// AllowedMachineTransitions defines valid transitions between machine states.
var AllowedMachineTransitions = map[MachineStatus][]MachineStatus{
	MachineStatusAvailable: {
		MachineStatusReserved,
		MachineStatusInUse,
		MachineStatusMaintenance,
		MachineStatusOffline,
	},
	MachineStatusReserved: {
		MachineStatusInUse,
		MachineStatusAvailable,
		MachineStatusOffline,
	},
	MachineStatusInUse: {
		MachineStatusAvailable,
		MachineStatusMaintenance,
		MachineStatusOffline,
	},
	MachineStatusMaintenance: {
		MachineStatusAvailable,
		MachineStatusOffline,
	},
	MachineStatusOffline: {
		MachineStatusAvailable,
	},
}

// Normalize ensures consistent uppercase status values.
func (s MachineStatus) Normalize() MachineStatus {
	return MachineStatus(strings.ToUpper(string(s)))
}

// CanTransitionTo checks if the current status can change to the given one.
func (s MachineStatus) CanTransitionTo(next MachineStatus) bool {
	next = next.Normalize()

	allowedNext, ok := AllowedMachineTransitions[s]
	if !ok {
		return false
	}

	return slices.Contains(allowedNext, next)
}

// AllowedNextStatuses returns all valid next statuses from the current one.
func (s MachineStatus) AllowedNextStatuses() []MachineStatus {
	return AllowedMachineTransitions[s.Normalize()]
}
