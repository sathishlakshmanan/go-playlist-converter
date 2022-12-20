package spotify

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
)

// structure of POST data to create new playlist
type PlaylistDetails struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}

func getUserInput() (string, PlaylistDetails) {

	scanner := bufio.NewScanner(os.Stdin)
	err := godotenv.Load()
	CheckError(err)

	userId := os.Getenv("USER_ID")
	/*
	 *uncomment the following lines to provide userId dynamically
	 *fmt.Println("Enter user id: ")
	 *scanner.Scan()
	 *userId := scanner.Text()
	 */

	endpoint := "https://api.spotify.com/v1/users/" + userId + "/playlists"

	fmt.Println("Enter a name for the playlist: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Enter a description for the playlist: ")
	scanner.Scan()
	description := scanner.Text()

	fmt.Println("Would you like your playlist to be public?")
	public := yesNo()

	details := PlaylistDetails{
		Name:        name,
		Description: description,
		Public:      public,
	}

	return endpoint, details
}

// yes or no prompt to set playlist created as public or private
func yesNo() bool {
	prompt := promptui.Select{
		Label: "Select[Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

type PlaylistResponse struct {
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	ID string `json:"id"`
}

func CreatePlaylist(client *http.Client, method string) (string, string) {

	endpoint, details := getUserInput()

	jsonData, err := json.Marshal(details)
	CheckError(err)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", os.Getenv("BEARER_TOKEN"))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := client.Do(req)
	CheckError(err)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	CheckError(err)

	var playlistResponse PlaylistResponse

	errr := json.Unmarshal(body, &playlistResponse)
	CheckError(errr)

	playlistID := fmt.Sprintf("%s", playlistResponse.ID)
	externalURL := fmt.Sprintf("%s", playlistResponse.ExternalUrls.Spotify)

	return playlistID, externalURL
}
