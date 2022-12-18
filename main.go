package main

import (
	"fmt"
	"regexp"

	//"go-playlist-converter/spotify"
	"go-playlist-converter/youtube"
	"net/http"
	"time"
)

func main() {
	//response := spotify.CreatePlaylist(httpClient(), http.MethodPost)
	//fmt.Println(string(response))

	ytResponse := youtube.GetPlaylist(httpClient(), http.MethodGet)

	for _, resp := range ytResponse.Items {

		re := regexp.MustCompile("[sS]ong[s |]|[vV]ideo[s |]|[|]")
		title := resp.Snippet.Title
		result := re.Split(title, 2)
		fmt.Println(result[0])
	}
}

// client to make requests
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}
