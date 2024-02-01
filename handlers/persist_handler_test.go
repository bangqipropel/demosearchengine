package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

// TestPersistMetadata tests the PersistMetadata function.
func TestPersistMetadata(t *testing.T) {
	// Create a new sync.Map instance for testing
	database := sync.Map{}

	// Construct a valid test payload
	validPayload := `title: Test App
version: 1.0.0
maintainers:
  - name: John Doe
    email: johndoe@example.com
company: Test Company
website: https://example.com
source: https://github.com/example/repo
license: Apache-2.0
description: Test description`

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/metadata", strings.NewReader(validPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		PersistMetadata(&database, w, r)
	})

	// Call the handler function, passing in the ResponseRecorder and the HTTP request
	handler.ServeHTTP(rr, req)

	// Check if the status code is http.StatusCreated
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Additional assertions can be added to check the response body or the metadata stored in the database

	// Check that using a non-POST method should return an error
	req, _ = http.NewRequest("GET", "/metadata", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code for GET request: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}
