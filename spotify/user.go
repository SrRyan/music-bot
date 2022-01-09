package spotify

import (
	"context"
)

type User struct {
	DisplayName  string       `json:"display_name"`
	ExternalURLs ExternalURLs `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Images       []Image      `json:"images"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// type PrivateData struct {
// 	Country     string `json:"country"`
// 	DisplayName string `json:"display_name"`
// 	Email       string `json:"email"`
// 	Product     string `json:"product"`
// 	Birthdate   string `json:"birthdate"`
// }

// 200, "OK"
// 401, "Unauthorized"
// 403, "Forbidden"
// 429, "Rate limit exceeded"

// GetCurrentUsersProfile returns the current user's public profile.
func (c *Client) GetCurrentUsersProfile(ctx context.Context) (*User, error) {

	user := new(User)

	err := c.Get(ctx, c.baseURL+"/me", user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
