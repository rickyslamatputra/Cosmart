package repository

import "Cosmart/internal/model"

// mockgen -source=C:\Users\ky\OneDrive\Desktop\Cosmart\internal\repository\repository.go -destination=C:\Users\ky\OneDrive\Desktop\Cosmart\internal/mock/mock_repository.go -package=mock

type PickupRepository interface {
	SaveSchedule(schedule model.PickupSchedule) error
	GetAllSchedule() []model.PickupSchedule
}

type BookRepository interface {
	GetBooksBySubject(subject string) ([]model.Book, error)
}
