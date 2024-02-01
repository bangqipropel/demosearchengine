package handlers

import (
	"demosearchengine/validators"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"

	"demosearchengine/metadata"
)

// PersistMetadata handles the storage of metadata. It only allows POST method.
func PersistMetadata(database *sync.Map, w http.ResponseWriter, r *http.Request) {
	// Check if the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse metadata from request body
	mPointer, err := metadata.NewMetadata(r.Body)
	metadata := *mPointer
	if err != nil {
		http.Error(w, "Error reading payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Add validation logic
	// Check if all required fields are present
	if !validators.ContainsRequiredFields(metadata) {
		http.Error(w, "Title, Company, Website, Maintainers, Source and License are required fields", http.StatusBadRequest)
		return
	}

	// Check if version is in the correct format
	if !validators.IsValidVersion(metadata.Version) {
		http.Error(w, "Version must be in the format of x.y.z", http.StatusBadRequest)
		return
	}

	// Check if maintainers are valid
	if !validators.ContainsValidMaintainers(metadata) {
		http.Error(w, "Each maintainer requires a name and a valid email", http.StatusBadRequest)
		return
	}

	// Generate a new UUID for the metadata
	metadata.ID = uuid.New().String()

	// Store the metadata object with the new ID in the database
	database.Store(metadata.ID, metadata) // Use Store method to add metadata

	// Send success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Metadata persisted successfully with ID: %s\n", metadata.ID)
}
