package types

type MachineStatus string

const (
	MachineStatusAvailable   MachineStatus = "available"
	MachineStatusInUse       MachineStatus = "in_use"
	MachineStatusMaintenance MachineStatus = "maintenance"
	MachineStatusOffline     MachineStatus = "offline"
	MachineStatusReserved    MachineStatus = "reserved"
)
