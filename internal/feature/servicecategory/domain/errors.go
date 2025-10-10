package domain

import "errors"

var (
	ErrServiceCategoryAlreadyExists = errors.New("service category already exists")
	ErrServiceCategoryNotFound      = errors.New("service category not found")
	ErrUnauthorizedAccess           = errors.New("unauthorized access to service category")
	ErrServiceCategoryDeleted       = errors.New("service category has been deleted")
)
