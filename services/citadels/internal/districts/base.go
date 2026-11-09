package districts

import (
	"os"
	"path/filepath"

	"github.com/Zadigo/citadels/internal/models"
	"gopkg.in/yaml.v3"
)

const (
	Nobel     = "Nobel"
	Religious = "Religious"
	Trade     = "Trade"
	Military  = "Military"
	Unique    = "Unique"
)

const UniqueDistrictCount int = 54

type District struct {
	models.BaseDistrict
}

func (d *District) GetId() string {
	return d.ID
}

func (d *District) GetPlayer() models.PlayerInterface {
	return d.Player
}

func (d *District) SetPlayer(player models.PlayerInterface) {
	d.Player = player
}

func (d *District) Buy(player models.PlayerInterface) error {
	return nil
}

type DistrictYaml struct {
	Districts struct {
		Unique []District
	}
}

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
