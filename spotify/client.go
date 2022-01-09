package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	http           *http.Client
	baseURL        string
	acceptLanguage string
}

func New(httpClient *http.Client) *Client {
	return &Client{
		http:    httpClient,
		baseURL: "https://api.spotify.com/v1",
	}
}

// GET wrapper for http.Get
func (c *Client) Get(ctx context.Context, url string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Handle response if not 200 OK
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

// GetSessionToken takes the clients cookie and returns a token.
// func (c *Client) GetToken(cookie *http.Cookie) (*oauth2.Token, error) {
// 	// Create a new token object, passing in the session cookie.

// 	// Return the token.
// 	return token, nil
// }
