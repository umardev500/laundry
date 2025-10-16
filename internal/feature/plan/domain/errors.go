package domain

import "fmt"

var (
	ErrPlanAlreadyExists      = fmt.Errorf("plan already exists")
	ErrPlanNotFound           = fmt.Errorf("plan not found")
	ErrPlanDeleted            = fmt.Errorf("plan has been deleted")
	ErrUnauthorizedPlanAccess = fmt.Errorf("unauthorized access to plan")
)
