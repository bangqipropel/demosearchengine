package metadata

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
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

// NewMetadata parses a Metadata instance from an io.Reader.
func NewMetadata(reader io.Reader) (*Metadata, error) {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var metadata Metadata
	err = yaml.Unmarshal(body, &metadata)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}
