package repository

import "Cosmart/internal/model"

// mockgen -source=C:\Users\ky\OneDrive\Desktop\Cosmart\internal\service\service.go -destination=C:\Users\ky\OneDrive\Desktop\Cosmart\internal/mock/mock_all.go -package=mock

type PickupService interface {
	SaveSchedule(schedule model.PickupSchedule) error
	GetAllSchedule() []model.PickupSchedule
}

type BookService interface {
	GetBooks(subject string) ([]model.Book, error)
}
