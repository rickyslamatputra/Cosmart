package pickup

import (
	"Cosmart/internal/model"
	"testing"
	"time"
)

func TestPickupRepository_SaveSchedule(t *testing.T) {
	// Arrange: create a new in-memory repository and a sample pickup schedule
	repo := &PickupRepository{}
	sampleSchedule := model.PickupSchedule{
		PickupTime: time.Now().Add(2 * time.Hour),
	}

	// Act: save the schedule
	err := repo.SaveSchedule(sampleSchedule)

	// Assert: ensure no error and that the schedule was saved correctly
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(repo.schedules) != 1 {
		t.Errorf("expected 1 schedule, got %d", len(repo.schedules))
	}

	if repo.schedules[0] != sampleSchedule {
		t.Errorf("expected schedule %+v, got %+v", sampleSchedule, repo.schedules[0])
	}
}

func TestPickupRepository_GetAllSchedule(t *testing.T) {
	// Arrange: create a new in-memory repository and add sample pickup schedules
	repo := &PickupRepository{}
	schedule1 := model.PickupSchedule{
		PickupTime: time.Now().Add(2 * time.Hour),
	}
	schedule2 := model.PickupSchedule{
		PickupTime: time.Now().Add(4 * time.Hour),
	}

	_ = repo.SaveSchedule(schedule1)
	_ = repo.SaveSchedule(schedule2)

	// Act: retrieve all schedules
	schedules := repo.GetAllSchedule()

	// Assert: ensure the correct number and content of schedules
	if len(schedules) != 2 {
		t.Errorf("expected 2 schedules, got %d", len(schedules))
	}

	if schedules[0] != schedule1 {
		t.Errorf("expected schedule1 %+v, got %+v", schedule1, schedules[0])
	}

	if schedules[1] != schedule2 {
		t.Errorf("expected schedule2 %+v, got %+v", schedule2, schedules[1])
	}
}
