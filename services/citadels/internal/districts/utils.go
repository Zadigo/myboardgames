package districts

import (
	"os"
	"path/filepath"

	"github.com/stretchr/testify/assert/yaml"
)

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

func CreateDistricts() {

}
