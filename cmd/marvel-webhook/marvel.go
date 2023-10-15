package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type MarvelResponse struct {
	Data struct {
		Results []struct {
			Name string `json:"name"`
		} `json:"results"`
	} `json:"data"`
}

// sanitizeLabel replaces spaces with underscores and removes illegal characters
func sanitizeLabel(label string) string {
	// Replace spaces with underscores
	label = strings.ReplaceAll(label, " ", "_")

	// Remove illegal characters using regular expression
	re := regexp.MustCompile(`[^a-zA-Z0-9_.-]`)
	label = re.ReplaceAllString(label, "")

	return label
}

// Function to construct the Marvel API URL
func constructMarvelAPIURL(ts string, publicKey string, hash string) string {
	return fmt.Sprintf("https://gateway.marvel.com/v1/public/characters?ts=%s&apikey=%s&hash=%s", ts, publicKey, hash)
}

func getRandomMarvelName(apiURL string, maxOffset int) (string, error) {
	var allNames []string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for offset := 0; offset <= maxOffset; offset += 100 {
		// Update the API URL with the new offset
		apiURLWithOffset := fmt.Sprintf("%s&offset=%d", apiURL, offset)

		resp, err := http.Get(apiURLWithOffset)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		var marvelResp MarvelResponse
		if err := json.Unmarshal(body, &marvelResp); err != nil {
			return "", err
		}

		for _, result := range marvelResp.Data.Results {
			allNames = append(allNames, result.Name)
		}
	}

	randomIndex := r.Intn(len(allNames))
	return allNames[randomIndex], nil
}
