package handlers

import (
	"demosearchengine/metadata"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// TestSearchMetadata tests the SearchMetadata function.
func TestSearchMetadata(t *testing.T) {
	// Initialize a new sync.Map instance for testing
	database := sync.Map{}

	// Create and store a sample metadata instance for testing
	sampleMetadata := metadata.Metadata{
		ID:      "123",
		Title:   "Test App",
		Version: "1.0.0",
		Maintainers: []struct {
			Name  string `yaml:"name"`
			Email string `yaml:"email"`
		}{
			{Name: "John Doe", Email: "johndoe@example.com"},
			{Name: "San Zhang ", Email: "sanzhang@example.com"},
		},
		Company:     "Test Company",
		Website:     "https://example.com",
		Source:      "https://github.com/example/repo",
		License:     "Apache-2.0",
		Description: "This is a test application",
	}
	database.Store(sampleMetadata.ID, sampleMetadata)

	// Create a new HTTP GET request with query parameters
	req, err := http.NewRequest("GET", "/search?id=123", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SearchMetadata(&database, w, r)
	})

	// Call the handler function, passing in the ResponseRecorder and the HTTP request
	handler.ServeHTTP(rr, req)

	// Check if the status code is http.StatusOK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var metadataObj metadata.Metadata
	err = json.Unmarshal(rr.Body.Bytes(), &metadataObj)
	if err != nil {
		t.Fatal(err)
	}
	if metadataObj.ID != sampleMetadata.ID {
		t.Errorf("handler returned unexpected body: got ID %v want ID %v", metadataObj.ID, sampleMetadata.ID)
	}
}
