package districts

import (
	"os"
	"path/filepath"

	"github.com/stretchr/testify/assert/yaml"
)

// ReadYamlFile reads the districts YAML file that contains the list
// of unique districts and basic ones that will be used in the game
func ReadYamlFile() (districts *DistrictYaml, err error) {
	absPath, err := filepath.Abs("../data/districts.yaml")
	if err != nil {
		return nil, err
	}

	if fileContent, err := os.ReadFile(absPath); err != nil {
		return nil, err
	} else {
		if err := yaml.Unmarshal(fileContent, &districts); err != nil {
			return nil, err
		} else {
			return districts, nil
		}
	}
}
