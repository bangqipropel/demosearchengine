package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"demosearchengine/metadata"
)

// SearchMetadata handles the search of metadata.
func SearchMetadata(database *sync.Map, w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	queryParams := r.URL.Query()
	idQuery := queryParams.Get("id")
	baseQuery := queryParams.Get("base_query")
	versionQuery := queryParams.Get("version")
	companyQuery := queryParams.Get("company")
	websiteQuery := queryParams.Get("website")
	licenseQuery := queryParams.Get("license")

	// If an ID is provided, prioritize search by ID
	if idQuery != "" {
		if value, ok := database.Load(idQuery); ok {
			metadata, ok := value.(metadata.Metadata)
			if !ok {
				http.Error(w, "Metadata found but type assertion failed", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(metadata)
			return
		} else {
			http.Error(w, "Metadata not found", http.StatusNotFound)
			return
		}
	}

	// Create a slice to save the search results
	var results []metadata.Metadata

	// Iterate through the database looking for matching metadata
	database.Range(func(key, value interface{}) bool {
		metadata, ok := value.(metadata.Metadata)
		if !ok {
			return true // Continue iterating
		}

		// Check if baseQuery matches in title, company, or description
		matchesBaseQuery := baseQuery == "" ||
			strings.Contains(strings.ToLower(metadata.Title), strings.ToLower(baseQuery)) ||
			strings.Contains(strings.ToLower(metadata.Company), strings.ToLower(baseQuery)) ||
			strings.Contains(strings.ToLower(metadata.Description), strings.ToLower(baseQuery))

		// Check other query parameters
		matchesVersion := versionQuery == "" || metadata.Version == versionQuery
		matchesCompany := companyQuery == "" || metadata.Company == companyQuery
		matchesWebsite := websiteQuery == "" || metadata.Website == websiteQuery
		matchesLicense := licenseQuery == "" || metadata.License == licenseQuery

		if matchesBaseQuery && matchesVersion && matchesCompany && matchesWebsite && matchesLicense {
			results = append(results, metadata)
		}

		return true // Continue iterating
	})

	// Encode the search results into JSON and send back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
