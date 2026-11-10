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
	PurpleSymbol string = "purple" // Unique
)

type DistrictInterface interface {
	GetId() string
	GetPlayer() PlayerInterface
	SetPlayer(player PlayerInterface)
	Buy(player PlayerInterface) error
	GetCost() int
	GetDistrictType() string
	GetSymbol() string
	IsUniqueDistrict() bool
	Resolve() error
}

type BaseDistrict struct {
	ID           string          `yaml:"id" json:"id"`
	Name         string          `yaml:"name" json:"name"`
	NameFr       string          `yaml:"name_fr" json:"name_fr"`
	Cost         int             `yaml:"cost" json:"cost"`
	Level        int             `yaml:"level" json:"level"`
	DistrictType string          `yaml:"district_type" json:"district_type"`
	Player       PlayerInterface `yaml:"player" json:"player"`
}

func (b *BaseDistrict) GetCost() int {
	return b.Cost
}

func (b *BaseDistrict) GetDistrictType() string {
	return b.DistrictType
}

func (b *BaseDistrict) GetSymbol() string {
	switch b.DistrictType {
	case TradeDistrict:
		return GreenSymbol
	case MilitaryDistrict:
		return RedSymbol
	case NobleDistrict:
		return BlueSymbol
	case ReligeousDistrict:
		return YellowSymbol
	case UniqueDistrict:
		return PurpleSymbol
	}
	return ""
}

func (b *BaseDistrict) GetId() string {
	return b.ID
}

func (b *BaseDistrict) IsUniqueDistrict() bool {
	return b.DistrictType == UniqueDistrict
}

func (b *BaseDistrict) Resolve() error {
	return nil
}
