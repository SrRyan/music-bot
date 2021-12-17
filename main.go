package main

func main() {
	println("Hello, World!")
}

// func NewHttpClient() {
// 	client := &http.Client{}
// 	req, err := http.NewRequest("GET", "http://example.com", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// }

// func NewSpotifyClient() {
// 	client := spotify.NewClient(nil, nil)
// 	client.Authenticate("", "", "", "")
// }
