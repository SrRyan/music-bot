package main

import (
	"net/http"
	"os"

	"github.com/SrRyan/music-bot/spotify"
	"github.com/joho/godotenv"
)

/*
- Add Auth
- Add AuthOptions
- Add Auth struct
- Add Spotify client
*/
func main() {

	println("Loading .env file...")
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	println("Loading .env file... Done")

	println("Creating Spotify client...")

	println("Creating Spotify Auth client...")
	auth := spotify.NewAuth(spotify.AuthOptions{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("SPOTIFY_REDIRECT_URL"),
	})
	println("Creating Spotify Auth client... Done")

	url := auth.GetAuthURL("1")

	println("Auth URL: " + url)

	http.Handle("/", http.FileServer(http.Dir("public")))
	// Handle the redirect from the spotify auth server
	http.HandleFunc("/callback", handleCallback)

	println("Starting server...")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
	println("Starting server... Done")

}

func handleCallback(w http.ResponseWriter, r *http.Request) {

	// TODO handle error and redirect to error page
	// https://developer.spotify.com/documentation/general/guides/authorization/code-flow/
	w.Write([]byte("Success! You can close this window now."))
}
