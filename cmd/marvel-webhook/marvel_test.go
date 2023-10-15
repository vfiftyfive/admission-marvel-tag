package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestGetRandomMarvelName(t *testing.T) {
	// Generate Marvel API URL dynamically
	ts := fmt.Sprintf("%v", time.Now().Unix())
	publicKey := "84f75d5854abb64040a580afa56dd9c0"
	privateKey := os.Getenv("MARVEL_PRIVATE_KEY")
	if privateKey == "" {
		t.Fatal("MARVEL_PRIVATE_KEY environment variable not set")
	}
	hash := fmt.Sprintf("%x", md5.Sum([]byte(ts+privateKey+publicKey)))
	apiURL := constructMarvelAPIURL(ts, publicKey, hash)

	log.Printf("Timestamp: %s, Hash: %s", ts, hash)
	log.Printf("Generated API URL: %s", apiURL)

	// Fetch and log the full API response for debugging
	resp, err := http.Get(apiURL)
	if err != nil {
		t.Fatalf("Error fetching API response: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading API response: %v", err)
	}

	log.Printf("Full API Response: %s", string(body))

	// Unmarshal and proceed with the test
	var marvelResp MarvelResponse
	if err := json.Unmarshal(body, &marvelResp); err != nil {
		t.Fatalf("Error unmarshaling API response: %v", err)
	}

	if len(marvelResp.Data.Results) == 0 {
		t.Fatalf("Error fetching Marvel name: no Marvel characters found")
	}

	_, err = getRandomMarvelName(apiURL, 100)
	if err != nil {
		t.Fatalf("Error fetching Marvel name: %v", err)
	}
}
