package domain

import (
	"encoding/json"
	"fmt"
)

type PlanFeatures struct {
	MaxUsers *int `json:"max_users"`
}

// ToMap converts a PlanFeatures struct into a map for Ent JSON storage.
func (f *PlanFeatures) ToMap() (map[string]any, error) {

	if f == nil {
		return nil, fmt.Errorf("plan features is nil")
	}

	// Marchat struct to JSON
	data, err := json.Marshal(f)
	if err != nil {
		return nil, fmt.Errorf("marshal plan features: %w", err)
	}

	// Unmarshal JSON into map
	var m map[string]any
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, fmt.Errorf("unmarshal plan features: %w", err)
	}

	return m, nil
}
