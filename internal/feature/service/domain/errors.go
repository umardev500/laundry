package domain

import "fmt"

var (
	ErrServiceAlreadyExists      = fmt.Errorf("service already exists")
	ErrServiceNotFound           = fmt.Errorf("service not found")
	ErrServiceDeleted            = fmt.Errorf("service has been deleted")
	ErrUnauthorizedServiceAccess = fmt.Errorf("unauthorized access to service")
)
