package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/SrRyan/music-bot/spotify"
	"github.com/joho/godotenv"
)

type global struct {
	mutex   sync.Mutex
	Auth    *spotify.Authentication
	Spotify *spotify.Client
}

var Global = global{}

func init() {
	println("Loading .env file...")
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	println("Loading .env file... Done")
}

func main() {

	println("Creating Spotify Auth client...")
	authentication := spotify.NewAuth(spotify.AuthOptions{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("SPOTIFY_REDIRECT_URL"),
	})
	println("Creating Spotify Auth client... Done")

	println("Creating Spotify client...")

	println("Creating Spotify client... Done")

	// TEMP Solution to get the auth code
	Global.Auth = authentication

	http.Handle("/", http.FileServer(http.Dir("public")))
	// Handle the redirect from the spotify auth server
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/user", handleGetUser)

	println("Starting server...")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
	println("Starting server... Done")

}

func handleGetUser(w http.ResponseWriter, r *http.Request) {

	// Check if the user is authenticated

	// Get the user
	user, err := Global.Spotify.GetCurrentUsersProfile(r.Context())
	if err != nil {
		fmt.Println("Error getting user: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	// Print the user
	fmt.Println(user)
}

// handleLogin handles the login page
func handleLogin(w http.ResponseWriter, r *http.Request) {
	state := generateState(w)
	http.Redirect(w, r, Global.Auth.GetAuthURL(state), http.StatusFound)
}

// generateState generates a random string to be used as a state
func generateState(w http.ResponseWriter) string {
	expiry := time.Now().Add(time.Hour * 24 * 7)

	b := make([]byte, 32)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "state", Value: state, Expires: expiry, HttpOnly: true}
	http.SetCookie(w, &cookie)

	return state
}

func handleCallback(w http.ResponseWriter, r *http.Request) {

	// Check if the state is valid
	cookie, err := r.Cookie("state")
	if err != nil {
		fmt.Println("Error getting cookie: ", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if r.FormValue("state") != cookie.Value {
		fmt.Println("State mismatch")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Exchange the code for an access token
	token, err := Global.Auth.ExchangeCodeForToken(r.Context(), r)
	if err != nil {
		fmt.Println("Error exchanging code for token: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	// Create a new spotify client
	spotifyClient := spotify.New(Global.Auth.HttpClient(r.Context(), token))

	// Get the user
	user, err := spotifyClient.GetCurrentUsersProfile(r.Context())
	if err != nil {
		fmt.Println("Error getting user: ", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	// Print the user
	fmt.Println(user)

	return

}

// func authorize(code string) (*oauth2.Token, error) {

// 	token, err := Global.Auth.ExchangeCodeForToken(code)
// 	if err != nil {
// 		return &oauth2.Token{}, err
// 	}

// 	fmt.Printf("Token Expires At: %s\n", token.Expiry)

// 	return token, nil
// }
