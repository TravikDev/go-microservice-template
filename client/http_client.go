package client

import "net/http"

// HTTP returns an HTTP client.
func HTTP() *http.Client {
	return &http.Client{}
}
