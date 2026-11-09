package models

import (
	"uuid"
)

// On his turn, a player can:
//   - Gather resources: They can either take two gold from the bank or draw
//     two district cards from the deck, choose one to keep, and discard the other
//     facedown to the bottom of the deck.
type PlayerChoices struct {
	GatherResources  bool
	BuildDistrict    bool
	SelectCharacter  bool
	TakeTwoGolds     bool
	DrawTwoDistricts bool
	DistrictToBuild  DistrictInterface
	CurrentGame      GameInterface
}

type PlayerInterface interface {
	GetId() string
	ApplyChoice(choices PlayerChoices) error
	GoldCount() int
	IncreaseGold(by int)
	DecreaseGold(by int)
	ResetGold()
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

func (p *Player) ApplyChoice(choice PlayerChoices) error {
	if choice.GatherResources {
		if !choice.TakeTwoGolds && !choice.DrawTwoDistricts {
			return PlayerChoiceError
		}

		if choice.TakeTwoGolds {
			p.IncreaseGold(2)
		} else if choice.DrawTwoDistricts {
			// Do soemthing
		}
	}

	if choice.BuildDistrict {
		if choice.DistrictToBuild == nil {
			return NoDistrictToBuild
		}
		// choice.DistrictToBuild.SetPlayer(p)
		return nil
	}

	return ChoiceDoesNotExistsError
}

// IncreaseGold adds gold to the player
func (p *Player) IncreaseGold(by int) {
	p.Gold += by
}

// DecreaseGold removes gold from the player, not below zero
func (p *Player) DecreaseGold(by int) {
	p.Gold -= by
	if p.Gold < 0 {
		p.Gold = 0
	}
}

func (p *Player) ResetGold() {
	p.Gold = 2
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
	for _, d := range p.Districts {
		d.SetPlayer(p)
	}
}

// TODO: Temporary location for player creation
func NewPlayer() PlayerInterface {
	return &Player{
		ID:                uuid.New().String(),
		Gold:              2,
		Districts:         []DistrictInterface{},
		SelectedCharacter: nil,
	}
}
