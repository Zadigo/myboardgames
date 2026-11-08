package models

import (
	"uuid"
)

type PlayerInterface interface {
	GetId() string
	GoldCount() int
	IncreaseGold(by int)
	DecreaseGold(by int)
	GetSelectedCharacter() CharacterInterface
	SetSelectedCharacter(c CharacterInterface)
	GetDistricts() []DistrictInterface
	AddDistrict(district DistrictInterface)
	ReplaceDistricts(districts []DistrictInterface)
}

type Player struct {
	ID                string
	Districts         []DistrictInterface
	Gold              int
	SelectedCharacter CharacterInterface
}

func (p *Player) GetId() string {
	return p.ID
}

func (p *Player) IncreaseGold(by int) {
	p.Gold += by
}

func (p *Player) DecreaseGold(by int) {
	p.Gold -= by
	if p.Gold < 0 {
		p.Gold = 0
	}
}

func (p *Player) GoldCount() int {
	return p.Gold
}

func (p *Player) GetSelectedCharacter() CharacterInterface {
	return p.SelectedCharacter
}

func (p *Player) SetSelectedCharacter(c CharacterInterface) {
	p.SelectedCharacter = c
}

func (p *Player) GetDistricts() []DistrictInterface {
	return p.Districts
}

// AddDistrict adds a new district to the player
func (p *Player) AddDistrict(district DistrictInterface) {
	// TODO: Check if the player was able to buy the district
	p.Districts = append(p.Districts, district)
}

// ReplaceDistricts replaces all districts of the player with new ones
func (p *Player) ReplaceDistricts(districts []DistrictInterface) {
	p.Districts = districts
}

// TODO: Temporary location for player creation
func NewPlayer() PlayerInterface {
	return &Player{
		ID:                uuid.New().String(),
		Gold:              0,
		Districts:         []DistrictInterface{},
		SelectedCharacter: nil,
	}
}
