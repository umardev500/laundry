package domain

import "fmt"

var (
	ErrMachineTypeAlreadyExists = fmt.Errorf("machine type already exists")
	ErrMachineTypeNotFound      = fmt.Errorf("machine type not found")
	ErrMachineTypeDeleted       = fmt.Errorf("machine type has been deleted")
)
