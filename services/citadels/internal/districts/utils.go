package districts

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"uuid"

	"github.com/stretchr/testify/assert/yaml"
)

type DistrictYamlData struct {
	Name   string `yaml:"name"`
	NameFr string `yaml:"name_fr"`
	Cost   int    `yaml:"cost"`
	Level  int    `yaml:"level"`
	Count  int    `yaml:"count"`
}

type DistrictYaml struct {
	Unique    []DistrictYamlData `yaml:"unique"`
	Religious []DistrictYamlData `yaml:"religious"`
	Military  []DistrictYamlData `yaml:"military"`
	Nobel     []DistrictYamlData `yaml:"nobel"`
	Trade     []DistrictYamlData `yaml:"trade"`
}

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

func GenerateDistricts() (districts []*District, err error) {
	// items := make([]*District, UniqueDistrictCount)
	items := []*District{}

	generator := func(dType string, d DistrictYamlData) {
		for range make([]struct{}, d.Count) {
			items = append(items, &District{
				ID:           uuid.New().String(),
				Name:         d.Name,
				Cost:         d.Cost,
				Player:       nil,
				NameFr:       d.NameFr,
				Level:        d.Level,
				DistrictType: dType,
			})
		}
	}

	if fileData, err := ReadYamlFile(); err == nil {
		for _, v := range fileData.Unique {
			generator(UniqueDistrict, v)
		}

		for _, v := range fileData.Religious {
			generator(ReligiousDistrict, v)
		}

		for _, v := range fileData.Military {
			generator(MilitaryDistrict, v)
		}

		for _, v := range fileData.Nobel {
			generator(NobelDistrict, v)
		}

		for _, v := range fileData.Trade {
			generator(TradeDistrict, v)
		}
	} else {
		return nil, errors.Join(fmt.Errorf("Could not generate districts"), err)
	}
	return items, nil
}

func ResolveUniqueDistrict(d *District) {
	if d.DistrictType == UniqueDistrict {
		switch d.Name {
		case "library":
			// Can keep all cards during the draw phase
		case "quarry":
			// Can build a district with the same name as existing ones
		case "haunted quarter":
			// Counts as a color of the player's choice when scoring
		case "keep":
			// Cannot be destroyed by the warlord (or by certain military effects)
		case "dragon gate":
			// Brings 8 victory points at the end the game
		case "school of magin":
			// Counts as a color of the player's choice when receiving gains from the player's character
		case "wishing well":
			// 1 point for every unique district including itself
		case "smithy":
			// Once per round, pay 2 golds to draw three cards
		default:
			// Unknown unique district
		}
	}
}
