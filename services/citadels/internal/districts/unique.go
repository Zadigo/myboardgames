package districts

import (
	"uuid"

	"github.com/Zadigo/citadels/internal/models"
)

func CreateUniqueDistricts() [14]models.BaseDistrict {
	districtNames := [14]models.BaseDistrict{
		{
			Name:  "Dragoon Gate",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Factory",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Haunted Quarter",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Imperial Treasury",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Keep",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Laboratory",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Library",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Map Room",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Quarry",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "School of Magic",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Smithy",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Statue",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Thieve's Den",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
		{
			Name:  "Wishing Well",
			Cost:  3,
			Color: "Red",
			Level: 1,
		},
	}

	for _, d := range districtNames {
		d.ID = uuid.New().String()
	}

	return districtNames
}
