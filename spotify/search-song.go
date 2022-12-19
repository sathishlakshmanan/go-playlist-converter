package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func SearchSong(client *http.Client, method string, query string) TrackResponse {

	endpoint := "https://api.spotify.com/v1/search"
	queryParams := map[string][]string{
		"q":     []string{query},
		"type":  []string{"track"},
		"limit": []string{"1"},
	}

	req, err := http.NewRequest(method, endpoint, nil)
	req.Header.Set("Authorization", os.Getenv("BEARER_TOKEN"))
	req.URL.RawQuery = url.Values(queryParams).Encode()

	response, err := client.Do(req)
	CheckError(err)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	CheckError(err)

	var songs TrackResponse
	errr := json.Unmarshal(body, &songs)
	CheckError(errr)

	return songs
}

type TrackResponse struct {
	Tracks struct {
		Items []struct {
			URI string `json:"uri"`
		} `json:"items"`
	} `json:"tracks"`
}
