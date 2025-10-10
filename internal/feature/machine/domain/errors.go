package domain

import "fmt"

var (
	ErrMachineAlreadyExists      = fmt.Errorf("machine already exists")
	ErrMachineNotFound           = fmt.Errorf("machine not found")
	ErrMachineDeleted            = fmt.Errorf("machine has been deleted")
	ErrUnauthorizedMachineAccess = fmt.Errorf("unauthorized access to machine")
)
