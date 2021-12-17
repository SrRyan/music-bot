package spotify

import "net/http"

type Client struct {
	httpClient *http.Client
}

// NewClient creates a new Spotify client
func NewClient(httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
	}
}
