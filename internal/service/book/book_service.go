package service

import (
	"Cosmart/internal/common"
	"Cosmart/internal/model"
	client "Cosmart/internal/pkg"
	"encoding/json"
	"fmt"
	"net/http"
)

type BookService struct {
	Client client.HTTPClient
}

func (s *BookService) GetBooks(subject string) ([]model.Book, error) {
	url := fmt.Sprintf(common.BaseURL, subject)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Make the HTTP request using the injected client
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch books, status code: %d", resp.StatusCode)
	}

	// Parse the response body
	var result struct {
		Books []struct {
			Title   string `json:"title"`
			Authors []struct {
				Name string `json:"name"`
			} `json:"authors"`
			EditionCount int `json:"edition_count"`
		} `json:"works"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// Convert the parsed result into a slice of Book
	books := make([]model.Book, len(result.Books))
	for i, work := range result.Books {
		authorName := ""
		if len(work.Authors) > 0 {
			authorName = work.Authors[0].Name
		}

		books[i] = model.Book{
			Title:   work.Title,
			Author:  authorName,
			Edition: work.EditionCount,
		}
	}

	return books, nil
}
