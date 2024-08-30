package handler

import (
	book "Cosmart/internal/service/book"
	"encoding/json"
	"net/http"
)

type BookHandler struct {
	Service *book.BookService
}

func (h *BookHandler) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	subject := r.URL.Query().Get("subject")
	if subject == "" {
		http.Error(w, "Subject is required", http.StatusBadRequest)
		return
	}

	books, err := h.Service.GetBooks(subject)
	if err != nil {
		http.Error(w, "Failed to fetch books: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
