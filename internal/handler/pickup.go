package handler

import (
	"Cosmart/internal/model"
	pickup "Cosmart/internal/service/pickup"
	"encoding/json"
	"net/http"
	"time"
)

type PickupHandler struct {
	Service *pickup.PickupService
}

func (h *PickupHandler) SchedulePickupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodGet:
		h.handleGet(w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *PickupHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var schedule model.PickupSchedule
	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if schedule.PickupTime.Before(time.Now()) {
		http.Error(w, "Pickup time must be in the future", http.StatusBadRequest)
		return
	}

	if err := s.Service.SchedulePickupService(schedule); err != nil {
		http.Error(w, "Failed to save schedule", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Schedule saved successfully"))
}

func (s *PickupHandler) handleGet(w http.ResponseWriter) {
	schedules := s.Service.GetAllSchedule()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(schedules); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
