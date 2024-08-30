package service

import (
	"Cosmart/internal/mock"
	"Cosmart/internal/model"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock.NewMockHTTPClient(ctrl)

	tests := []struct {
		name          string
		subject       string
		mockResponse  *http.Response
		mockError     error
		expectedBooks []model.Book
		expectedError error
	}{
		{
			name:    "Success - Books Found",
			subject: "fiction",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       createMockBody(`{"works":[{"title":"Book One","authors":[{"name":"Author One"}],"edition_count":1},{"title":"Book Two","authors":[{"name":"Author Two"}],"edition_count":2}]}`),
			},
			mockError: nil,
			expectedBooks: []model.Book{
				{Title: "Book One", Author: "Author One", Edition: 1},
				{Title: "Book Two", Author: "Author Two", Edition: 2},
			},
			expectedError: nil,
		},
		{
			name:    "Failure - Non-200 Status Code",
			subject: "non-fiction",
			mockResponse: &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       createMockBody(""),
			},
			mockError:     nil,
			expectedBooks: nil,
			expectedError: errors.New("failed to fetch books, status code: 404"),
		},
		{
			name:    "Failure - HTTP Request Error",
			subject: "error",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       createMockBody(""),
			},
			mockError:     errors.New("request error"),
			expectedBooks: nil,
			expectedError: errors.New("request error"),
		},
		{
			name:    "Failure - JSON Decode Error",
			subject: "json-error",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       createMockBody(`invalid json`),
			},
			mockError:     nil,
			expectedBooks: nil,
			expectedError: &json.SyntaxError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient.EXPECT().Do(gomock.Any()).Return(tt.mockResponse, tt.mockError).Times(1)

			service := &BookService{Client: mockClient}

			books, err := service.GetBooks(tt.subject)

			assert.ElementsMatch(t, tt.expectedBooks, books)

			if tt.expectedError != nil {
				assert.ErrorAs(t, err, &tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Helper function to create a mock body for HTTP responses
func createMockBody(content string) *MockReadCloser {
	return &MockReadCloser{
		ReadFunc: func(p []byte) (n int, err error) {
			copy(p, content)
			return len(content), nil
		},
		CloseFunc: func() error {
			return nil
		},
	}
}

// MockReadCloser is a mock implementation of io.ReadCloser
type MockReadCloser struct {
	ReadFunc  func(p []byte) (n int, err error)
	CloseFunc func() error
}

func (m *MockReadCloser) Read(p []byte) (n int, err error) {
	return m.ReadFunc(p)
}

func (m *MockReadCloser) Close() error {
	return m.CloseFunc()
}
