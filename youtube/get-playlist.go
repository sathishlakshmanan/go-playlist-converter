package youtube

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func getUserInput() []string {

	scanner := bufio.NewScanner(os.Stdin)
	err := godotenv.Load()
	CheckError(err)

	// get playlist id from user
	fmt.Println("Enter the Youtube playlist id: ")
	scanner.Scan()
	playlistID := []string{scanner.Text()}

	return playlistID
}

// function to log errors
func CheckError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// get playlist from youtube
func GetPlaylist(client *http.Client, method string) Fields {

	endpoint := "https://youtube.googleapis.com/youtube/v3/playlistItems"
	queryParams := map[string][]string{
		"playlistId": getUserInput(),
		"part":       []string{"snippet"},
		"key":        []string{os.Getenv("YOUTUBE_API_KEY")},
		"fields":     []string{"items(snippet)"},
		"maxResults": []string{"10"},
	}

	req, err := http.NewRequest(method, endpoint, nil)
	req.URL.RawQuery = url.Values(queryParams).Encode()

	response, err := client.Do(req)
	CheckError(err)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	CheckError(err)

	var songs Fields
	errr := json.Unmarshal(body, &songs)
	CheckError(errr)

	return songs
}

type Fields struct {
	Items []struct {
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
	} `json:"items"`
}
