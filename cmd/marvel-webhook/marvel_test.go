package main

import (
	"testing"
)

// TestGetRandomMarvelName tests the getRandomMarvelName function
func TestGetRandomMarvelName(t *testing.T) {
	// Call the function to get a random Marvel name
	marvelName, err := getRandomMarvelName()
	if err != nil {
		t.Fatalf("Failed to get random Marvel name: %v", err)
	}

	// Check if the Marvel name is not empty
	if marvelName == "" {
		t.Fatalf("Received empty Marvel name")
	}

	t.Logf("Received Marvel name: %s", marvelName)
}
