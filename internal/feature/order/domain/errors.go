package domain

import "fmt"

var (
	ErrGuestEmailOrPhoneRequired = fmt.Errorf("guest email or phone is required")
)
