package pickup

import "Cosmart/internal/model"

type PickupRepository struct {
	schedules []model.PickupSchedule
}

// NewOpenLibraryRepository creates a new instance of OpenLibraryRepository with an HTTP client.
func NewPickupRepository() *PickupRepository {
	return &PickupRepository{}
}

func (r *PickupRepository) SaveSchedule(schedule model.PickupSchedule) error {
	r.schedules = append(r.schedules, schedule)
	return nil
}

func (r *PickupRepository) GetAllSchedule() []model.PickupSchedule {
	return r.schedules
}
