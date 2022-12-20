package spotify

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func AddSongsToPlaylist(client *http.Client, playlistID string, songURIS string) []byte {

	endpoint := "https://api.spotify.com/v1/playlists/" + playlistID + "/tracks"

	queryParams := map[string][]string{
		"uris": []string{songURIS},
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	req.Header.Set("Authorization", os.Getenv("BEARER_TOKEN"))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.URL.RawQuery = url.Values(queryParams).Encode()
	CheckError(err)

	response, err := client.Do(req)
	CheckError(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	CheckError(err)

	return body
}
