package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sathishlakshmanan/go-playlist-converter/spotify"
	"github.com/sathishlakshmanan/go-playlist-converter/youtube"
	"net/http"
	"time"
)

func main() {

	client := httpClient()

	playlistID, externalURL := spotify.CreatePlaylist(client, http.MethodPost)

	ytResponse := youtube.GetPlaylist(client, http.MethodGet)

	var songs []string

	for _, resp := range ytResponse.Items {

		re := regexp.MustCompile("[sS]ong[s |]|[vV]ideo[s |]|[|]")
		title := resp.Snippet.Title
		result := re.Split(title, 2)
		song := spotify.SearchSong(client, http.MethodGet, result[0])
		songs = append(songs, song)
	}

	songsCombined := strings.Join(songs, ",")
	spotify.AddSongsToPlaylist(client, playlistID, songsCombined)
	fmt.Printf("Here is your playlist: %v", externalURL)
}

// client to make requests
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}
