package validators

import (
	"demosearchengine/metadata"
	"testing"
)

// TestIsValidEmail tests the IsValidEmail function.
func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected bool
	}{
		{"validemail@example.com", true},
		{"invalid-email", false},
		{"", false},
	}

	for _, tc := range testCases {
		if IsValidEmail(tc.email) != tc.expected {
			t.Errorf("IsValidEmail(%s) = %v, expected %v", tc.email, !tc.expected, tc.expected)
		}
	}
}

// TestIsValidVersion tests the IsValidVersion function.
func TestIsValidVersion(t *testing.T) {
	testCases := []struct {
		version  string
		expected bool
	}{
		{"1.0.0", true},
		{"0.1.0", true},
		{"1.0", false},
		{"invalid", false},
	}

	for _, tc := range testCases {
		if IsValidVersion(tc.version) != tc.expected {
			t.Errorf("IsValidVersion(%s) = %v, expected %v", tc.version, !tc.expected, tc.expected)
		}
	}
}

// TestContainsRequiredFields tests the ContainsRequiredFields function.
func TestContainsRequiredFields(t *testing.T) {
	validMetadata := metadata.Metadata{
		Title:   "Valid Title",
		Company: "Valid Company",
		Website: "https://validwebsite.com",
		Source:  "https://validsource.com",
		License: "Valid License",
		Maintainers: []struct {
			Name  string `yaml:"name"`
			Email string `yaml:"email"`
		}{{Name: "John Doe", Email: "john@example.com"}},
	}

	if !ContainsRequiredFields(validMetadata) {
		t.Error("ContainsRequiredFields should return true for valid metadata")
	}

	invalidMetadata := metadata.Metadata{
		// Missing required fields
	}

	if ContainsRequiredFields(invalidMetadata) {
		t.Error("ContainsRequiredFields should return false for invalid metadata")
	}
}

// TestContainsValidMaintainers tests the ContainsValidMaintainers function.
func TestContainsValidMaintainers(t *testing.T) {
	validMetadata := metadata.Metadata{
		Maintainers: []struct {
			Name  string `yaml:"name"`
			Email string `yaml:"email"`
		}{{Name: "John Doe", Email: "john@example.com"}},
	}

	if !ContainsValidMaintainers(validMetadata) {
		t.Error("ContainsValidMaintainers should return true for valid maintainers")
	}

	invalidMetadata := metadata.Metadata{
		Maintainers: []struct {
			Name  string `yaml:"name"`
			Email string `yaml:"email"`
		}{{Name: "", Email: "invalid-email"}},
	}

	if ContainsValidMaintainers(invalidMetadata) {
		t.Error("ContainsValidMaintainers should return false for invalid maintainers")
	}
}
