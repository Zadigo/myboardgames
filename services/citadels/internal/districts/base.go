package districts

import "github.com/Zadigo/citadels/internal/models"

const (
	Nobel     = "Nobel"
	Religious = "Religious"
	Trade     = "Trade"
	Military  = "Military"
	Unique    = "Unique"
)

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
