package models

const (
	ReligeousDistrict string = "religious"
	MilitaryDistrict  string = "military"
	NobleDistrict     string = "noble"
	TradeDistrict     string = "trade"
	UniqueDistrict    string = "unique"
)

const (
	RedSymbol    string = "red"    // Military
	GreenSymbol  string = "green"  // Trade
	BlueSymbol   string = "blue"   // Noble
	YellowSymbol string = "yellow" // Religious
)

type DistrictInterface interface {
	GetId() string
	GetPlayer() PlayerInterface
	SetPlayer(player PlayerInterface)
	Buy(player PlayerInterface) error
}

type BaseDistrict struct {
	ID     string          `yaml:"id" json:"id"`
	Name   string          `yaml:"name" json:"name"`
	NameFr string          `yaml:"name_fr" json:"name_fr"`
	Cost   int             `yaml:"cost" json:"cost"`
	Level  int             `yaml:"level" json:"level"`
	Symbol string          `yaml:"symbol" json:"symbol"`
	Player PlayerInterface `yaml:"player" json:"player"`
}
