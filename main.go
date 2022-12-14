package main

import (
	"fmt"
	"go-playlist-converter/spotify"
	"net/http"
	"time"
)

func main() {
	response := spotify.CreatePlaylist(httpClient(), http.MethodPost)
	fmt.Println(string(response))
}

// client to make requests
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}
