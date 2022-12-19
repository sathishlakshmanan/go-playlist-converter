package main

import (
	"fmt"
	"regexp"

	"go-playlist-converter/spotify"
	"go-playlist-converter/youtube"
	"net/http"
	"time"
)

func main() {

	client := httpClient()

	// response := spotify.CreatePlaylist(client, http.MethodPost)
	// fmt.Println(string(response))

	ytResponse := youtube.GetPlaylist(client, http.MethodGet)

	var songs []string

	for _, resp := range ytResponse.Items {

		re := regexp.MustCompile("[sS]ong[s |]|[vV]ideo[s |]|[|]")
		title := resp.Snippet.Title
		result := re.Split(title, 2)
		song := spotify.SearchSong(client, http.MethodGet, result[0])
		songs = append(songs, fmt.Sprintf("%+v", song.Tracks.Items))
	}
	fmt.Println(songs)
}

// client to make requests
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}
