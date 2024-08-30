package main

import (
	"Cosmart/internal/handler"
	pickupRepo "Cosmart/internal/repository/pickup"
	bookService "Cosmart/internal/service/book"
	pickupService "Cosmart/internal/service/pickup"
	"log"
	"net/http"
)

func main() {
	// Initialize the HTTP client
	httpClient := &http.Client{}

	// Initialize the repositories
	pickupRepo := pickupRepo.NewPickupRepository()

	// Initialize the services with the repositories
	bookService := &bookService.BookService{Client: httpClient}
	pickupService := &pickupService.PickupService{Repo: pickupRepo}

	// Initialize the handlers with the services
	bookHandler := &handler.BookHandler{Service: bookService}
	pickupHandler := &handler.PickupHandler{Service: pickupService}

	// Register handlers
	http.HandleFunc("/api/books", bookHandler.GetBooksHandler)                   // GET /api/books?subject=subjectName
	http.HandleFunc("/api/schedule-pickup", pickupHandler.SchedulePickupHandler) // POST /api/schedule-pickup
	http.HandleFunc("/api/get-schedule", pickupHandler.SchedulePickupHandler)    // GET /api/get-schedule

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
