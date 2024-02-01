package main

import (
	"fmt"
	"net/http"
	"sync"

	"demosearchengine/handlers"
)

// Metadata struct defines the expected structure for the API.
type Metadata struct {
	ID          string `json:"id"`
	Title       string `yaml:"title"`
	Version     string `yaml:"version"`
	Maintainers []struct {
		Name  string `yaml:"name"`
		Email string `yaml:"email"`
	} `yaml:"maintainers"`
	Company     string `yaml:"company"`
	Website     string `yaml:"website"`
	Source      string `yaml:"source"`
	License     string `yaml:"license"`
	Description string `yaml:"description"`
}

// Using sync.Map instead of a slice to improve thread safety.
var database sync.Map

func main() {
	// Initialize a sync.Map instance
	var database sync.Map

	// Create a closure handler function that passes the database to persistMetadata
	persistHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.PersistMetadata(&database, w, r)
	}

	// Create a closure handler function that passes the database to searchMetadata
	searchHandler := func(w http.ResponseWriter, r *http.Request) {
		handlers.SearchMetadata(&database, w, r)
	}

	http.HandleFunc("/metadata", persistHandler)
	http.HandleFunc("/search", searchHandler)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
