package service

import (
	"Cosmart/internal/mock"
	"Cosmart/internal/model"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestSchedulePickupService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		schedule      model.PickupSchedule
		mockRepo      *mock.MockPickupRepository
		expectedError error
	}{
		{
			name: "Success - Schedule Saved",
			schedule: model.PickupSchedule{
				PickupTime: time.Now().Add(1 * time.Hour),
			},
			mockRepo: func() *mock.MockPickupRepository {
				mr := mock.NewMockPickupRepository(ctrl)
				mr.EXPECT().SaveSchedule(gomock.Any()).Return(nil).Times(1)
				return mr
			}(),
			expectedError: nil,
		},
		{
			name: "Failure - Save Schedule Error",
			schedule: model.PickupSchedule{
				PickupTime: time.Now().Add(1 * time.Hour),
			},
			mockRepo: func() *mock.MockPickupRepository {
				mr := mock.NewMockPickupRepository(ctrl)
				mr.EXPECT().SaveSchedule(gomock.Any()).Return(errors.New("save error")).Times(1)
				return mr
			}(),
			expectedError: errors.New("save error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &PickupService{Repo: tt.mockRepo}

			err := service.SchedulePickupService(tt.schedule)

			if err != nil && tt.expectedError != nil {
				if err.Error() != tt.expectedError.Error() {
					t.Errorf("expected error %q, got %q", tt.expectedError.Error(), err.Error())
				}
			} else if (err == nil) != (tt.expectedError == nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}

func TestGetAllSchedule(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	referenceTime := time.Date(2024, time.August, 30, 10, 0, 0, 0, time.UTC)

	tests := []struct {
		name              string
		mockRepo          *mock.MockPickupRepository
		expectedSchedules []model.PickupSchedule
	}{
		{
			name: "Success - Get All Schedules",
			mockRepo: func() *mock.MockPickupRepository {
				mr := mock.NewMockPickupRepository(ctrl)
				mr.EXPECT().GetAllSchedule().Return([]model.PickupSchedule{
					{PickupTime: referenceTime.Add(2 * time.Hour)},
					{PickupTime: referenceTime.Add(4 * time.Hour)},
				}).Times(1)
				return mr
			}(),
			expectedSchedules: []model.PickupSchedule{
				{PickupTime: referenceTime.Add(2 * time.Hour)},
				{PickupTime: referenceTime.Add(4 * time.Hour)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &PickupService{Repo: tt.mockRepo}

			schedules := service.GetAllSchedule()

			if len(schedules) != len(tt.expectedSchedules) {
				t.Errorf("expected %v schedules, got %v", len(tt.expectedSchedules), len(schedules))
			}

			for i, schedule := range schedules {
				if schedule != tt.expectedSchedules[i] {
					t.Errorf("expected schedule %v, got %v", tt.expectedSchedules[i], schedule)
				}
			}
		})
	}
}
