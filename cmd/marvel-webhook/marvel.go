package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Structure for Marvel API response
type MarvelResponse struct {
	Data struct {
		Results []struct {
			Name string `json:"name"`
		} `json:"results"`
	} `json:"data"`
}

// Function to fetch a random Marvel character name
func getRandomMarvelName() (string, error) {
	// Make an HTTP GET request to the Marvel API
	resp, err := http.Get("https://gateway.marvel.com/v1/public/characters?YOUR_API_KEY_HERE")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Deserialize the JSON response
	var marvelResp MarvelResponse
	if err := json.Unmarshal(body, &marvelResp); err != nil {
		return "", err
	}

	// Pick a random Marvel character name
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Intn(len(marvelResp.Data.Results))
	return marvelResp.Data.Results[randomIndex].Name, nil
}
