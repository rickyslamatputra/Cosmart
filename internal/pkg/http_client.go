// http_client.go
package service

import (
	"net/http"
)

// mockgen -source=C:\Users\ky\OneDrive\Desktop\Cosmart\internal\pkg\http_client.go -destination=C:\Users\ky\OneDrive\Desktop\Cosmart\internal/mock/mock_http_client.go -package=mock

// HTTPClient is an interface for making HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
