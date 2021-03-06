package spotify

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	AuthURL  string = "https://accounts.spotify.com/authorize"
	TokenURL string = "https://accounts.spotify.com/api/token"
)

// Scopes is a set of scopes that can be used when requesting a new token.
const (

	// IMAGES
	// Write access to user-provided images.
	UgcImageUpload string = "ugc-image-upload"

	// SPOTIFY CONNECT
	// Read access to a user’s player state.
	UserReadPlaybackState string = "user-read-playback-state"
	// Write access to a user’s playback state
	UserModifyPlaybackState string = "user-follow-modify"
	// Read access to a user’s currently playing content.
	UserReadCurrentlyPlaying string = "user-read-currently-playing"

	// USERS
	// Read access to user’s subscription details (type of user account).
	UserReadPrivate string = "user-read-private"
	// Read access to user’s email address.
	UserReadEmail string = "user-read-email"

	// FOLLOW
	// Write/delete access to the list of artists and other users that the user follows.
	UserFollowModify string = "user-follow-modify"
	// Read access to the list of artists and other users that the user follows.
	UserFollowRead string = "user-follow-read"

	// LIBRARY
	// Write/delete access to a user's "Your Music" library.
	UserLibraryModify string = "user-library-modify"
	// Read access to a user's library.
	UserLibraryRead string = "user-library-read"

	// PLAYBACK
	// Control playback of a Spotify track. This scope is currently available to the Web Playback SDK. The user must have a Spotify Premium account.
	Streaming string = "streaming"
	// Remote control playback of Spotify. This scope is currently available to Spotify iOS and Android SDKs.
	AppRemoteControl string = "app-remote-control"

	// LISTENING HISTORY
	// Read access to a user’s playback position in a content.
	UserReadPlaybackPosition string = "user-read-playback-position"
	// Read access to a user's top artists and tracks.
	UserTopRead string = "user-top-read"
	// Read access to a user’s recently played tracks.
	UserReadRecentlyPlayed string = "user-read-recently-played"

	// PLAYLIST
	// Write/delete access to a user's private playlists.
	PlaylistModifyPrivate string = "playlist-modify-private"
	// Include collaborative playlists when requesting a user's playlists.
	PlaylistReadCollaborative string = "playlist-read-collaborative"
	// Read access to user's private playlists.
	PlaylistReadPrivate string = "playlist-read-private"
	// Write/delete access to a user's public playlists.
	PlaylistModifyPublic string = "playlist-modify-public"
)

type AuthOptions struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

type Authentication struct {
	Config *oauth2.Config
}

// NewAuth returns a new Auth struct
func NewAuth(opts AuthOptions) *Authentication {
	cfg := &oauth2.Config{
		ClientID:     opts.ClientID,
		ClientSecret: opts.ClientSecret,
		RedirectURL:  opts.RedirectURL,
		Scopes:       opts.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: TokenURL,
		},
	}

	return &Authentication{
		Config: cfg,
	}

}

// GetCode returns the code from the URL
func (a *Authentication) GetCode(r *http.Request) string {
	// TODO Handle Error if URL contains error or other then expected
	return r.URL.Query().Get("code")
}

// AuthURL returns the URL used for obtaining an access token.
func (a *Authentication) GetAuthURL(state string) string {
	return a.Config.AuthCodeURL(state)
}

// Exchange converts an authorization code into a token.
func (a *Authentication) ExchangeCodeForToken(ctx context.Context, req *http.Request, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {

	state := req.URL.Query().Get("state")
	if state == "" {
		return nil, errors.New("state mismatch")
	}

	// TODO Handle Error if URL contains error or other then expected
	code := req.URL.Query().Get("code")
	if code == "" {
		return nil, errors.New("missing code")
	}

	return a.Config.Exchange(ctx, code, opts...)

}

// HttpClient returns an http.Client that uses the provided token.
// The token will automatically be refreshed as necessary.
// The returned client and its Transport should not be modified.
func (a *Authentication) HttpClient(ctx context.Context, token *oauth2.Token) *http.Client {
	return a.Config.Client(ctx, token)
}
