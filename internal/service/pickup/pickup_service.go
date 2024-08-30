package service

import (
	"Cosmart/internal/model"
	repo "Cosmart/internal/repository"
)

type PickupService struct {
	Repo repo.PickupRepository
}

func (s *PickupService) SchedulePickupService(schedule model.PickupSchedule) error {
	if err := s.Repo.SaveSchedule(schedule); err != nil {
		return err
	}
	return nil
}

func (s *PickupService) GetAllSchedule() []model.PickupSchedule {
	return s.Repo.GetAllSchedule()
}
