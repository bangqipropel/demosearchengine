package validators

import (
	"net/mail"
	"regexp"

	"demosearchengine/metadata"
)

// IsValidEmail checks if the email address is valid.
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// IsValidVersion checks if the version number follows the semantic versioning format.
func IsValidVersion(version string) bool {
	re := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	return re.MatchString(version)
}

// ContainsRequiredFields checks if the Metadata contains all required fields.
// Required fields are Title, Company, Website, Source, and License.
func ContainsRequiredFields(metadata metadata.Metadata) bool {
	if metadata.Title == "" || metadata.Company == "" || metadata.Website == "" || metadata.Source == "" || metadata.License == "" {
		return false
	}
	if len(metadata.Maintainers) == 0 {
		return false
	}
	return true
}

// ContainsValidMaintainers checks if each maintainer in the Metadata has a name and a valid email.
func ContainsValidMaintainers(metadata metadata.Metadata) bool {
	for _, maintainer := range metadata.Maintainers {
		if maintainer.Name == "" || !IsValidEmail(maintainer.Email) {
			return false
		}
	}
	return true
}
