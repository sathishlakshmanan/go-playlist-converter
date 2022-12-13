package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {
	resp := sendRequest(httpClient(), http.MethodPost, "https://httpbin.org/post")
	fmt.Println(string(resp))
}

// client to make requests
func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

// structure of POST data to create new playlist
type PlaylistDetails struct {
	Name        string
	Description string
	Public      bool
}

func createPlaylist(userId string) PlaylistDetails {
	endpoint := "https://api.spotify.com/v1/users/" + userId + "/playlists"
	fmt.Println(endpoint)

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	name := scanner.Text()

	scanner.Scan()
	description := scanner.Text()

	public := yesNo()

	details := PlaylistDetails{
		Name:        name,
		Description: description,
		Public:      public,
	}

	return details
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

func sendRequest(client *http.Client, method string, endpoint string) []byte {
	details := createPlaylist("")

	jsonData, err := json.Marshal(details)

	CheckError(err)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := client.Do(req)

	CheckError(err)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	CheckError(err)

	return body
}
