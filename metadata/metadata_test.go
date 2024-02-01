package metadata

import (
	"strings"
	"testing"
)

// TestNewMetadata tests the NewMetadata function for correct parsing of YAML data.
func TestNewMetadata(t *testing.T) {
	// Sample YAML data for testing
	yamlData := `
id: "123"
title: "Test Title"
version: "1.0.0"
maintainers:
  - name: "John Doe"
    email: "johndoe@example.com"
  - name: "San Zhang"
    email: "sanzhang@example.com"
company: "Test Company"
website: "https://example.com"
source: "https://github.com/example/repo"
license: "Apache-2.0"
description: "This is a test description"`

	// Use strings.NewReader to create an io.Reader
	reader := strings.NewReader(yamlData)

	// Call NewMetadata and pass the reader
	metadata, err := NewMetadata(reader)
	if err != nil {
		t.Fatalf("NewMetadata failed: %v", err)
	}

	// Perform checks to see if metadata fields are correctly populated
	if metadata.ID != "123" {
		t.Errorf("Expected ID to be '123', got '%s'", metadata.ID)
	}
	if metadata.Title != "Test Title" {
		t.Errorf("Expected Title to be 'Test Title', got '%s'", metadata.Title)
	}
	if metadata.Version != "1.0.0" {
		t.Errorf("Expected Version to be '1.0.0', got '%s'", metadata.Version)
	}
	if len(metadata.Maintainers) != 2 ||
		metadata.Maintainers[0].Name != "John Doe" ||
		metadata.Maintainers[1].Name != "San Zhang" ||
		metadata.Maintainers[0].Email != "johndoe@example.com" ||
		metadata.Maintainers[1].Email != "sanzhang@example.com" {
		t.Errorf("Maintainers data did not match expected values")
	}
	if metadata.Company != "Test Company" {
		t.Errorf("Expected Company to be 'Test Company', got '%s'", metadata.Company)
	}
	if metadata.Website != "https://example.com" {
		t.Errorf("Expected Website to be 'https://example.com', got '%s'", metadata.Website)
	}
	if metadata.Source != "https://github.com/example/repo" {
		t.Errorf("Expected Source to be 'https://github.com/example/repo', got '%s'", metadata.Source)
	}
	if metadata.License != "Apache-2.0" {
		t.Errorf("Expected License to be 'Apache-2.0', got '%s'", metadata.License)
	}
	if metadata.Description != "This is a test description" {
		t.Errorf("Expected Description to be 'This is a test description', got '%s'", metadata.Description)
	}
}
