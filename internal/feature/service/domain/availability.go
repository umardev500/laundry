package domain

import "github.com/google/uuid"

type AvailabilityResult struct {
	RequestedIDs      []uuid.UUID
	AvailableServices []*Service
}

// AllAvailable returns true if all requested IDs exist.
func (r *AvailabilityResult) AllAvailable() bool {
	return len(r.AvailableServices) == len(r.RequestedIDs)
}

// UnavailableIDs returns a list of item IDs that are not available.
func (r *AvailabilityResult) UnavailableIDs() []uuid.UUID {
	// UnavailableIDs returns the list of requested IDs not found in AvailableServices.
	found := make(map[uuid.UUID]bool, len(r.AvailableServices))
	for _, s := range r.AvailableServices {
		found[s.ID] = true
	}

	var missing []uuid.UUID
	for _, id := range r.RequestedIDs {
		if !found[id] {
			missing = append(missing, id)
		}
	}
	return missing
}
