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

// MarvelResponse represents the response from the Marvel API
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

// Function to fetch a random Marvel name
func getRandomMarvelName(apiURL string, maxOffset int) (string, error) {
	var allNames []string

	// Initialize a new random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Iterate over the Marvel API results in increments of 100
	for offset := 0; offset <= maxOffset; offset += 100 {
		// Update the API URL with the new offset
		apiURLWithOffset := fmt.Sprintf("%s&offset=%d", apiURL, offset)

		// Fetch the API response
		resp, err := http.Get(apiURLWithOffset)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		// Read the API response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		// Unmarshal the API response
		var marvelResp MarvelResponse
		if err := json.Unmarshal(body, &marvelResp); err != nil {
			return "", err
		}

		// Append all Marvel character names to the allNames slice
		for _, result := range marvelResp.Data.Results {
			allNames = append(allNames, result.Name)
		}
	}

	// Pick a random Marvel name from the allNames slice
	randomIndex := r.Intn(len(allNames))
	return allNames[randomIndex], nil
}
