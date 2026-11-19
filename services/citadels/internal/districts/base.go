package districts

import (
	"github.com/Zadigo/citadels/internal/models"
)

const (
	NobelDistrict     string = "Nobel"
	ReligiousDistrict string = "Religious"
	TradeDistrict     string = "Trade"
	MilitaryDistrict  string = "Military"
	UniqueDistrict    string = "Unique"
)

const UniqueDistrictCount int = 84

type District struct {
	models.BaseDistrict
}

func (d *District) Buy(player models.PlayerInterface) error {
	return nil
}
