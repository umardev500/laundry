package types

type Status string

const (
	StatusActive    Status = "active"
	StatusSuspended Status = "suspended"
	StatusDeleted   Status = "deleted"
)

func (s Status) IsActive() bool {
	return s == StatusActive
}

func (s Status) IsSuspended() bool {
	return s == StatusSuspended
}

func (s Status) IsDeleted() bool {
	return s == StatusDeleted
}
