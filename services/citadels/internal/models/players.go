package models

import (
	"sync/atomic"
	"uuid"

	"github.com/gorilla/websocket"
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
	GetBuiltDistricts() []DistrictInterface
	GetCurrentHand() []DistrictInterface
	GetPoints() int
	SetSelectedCharacter(c CharacterInterface)
	BuildDistrict(district DistrictInterface)
	DrawDistricts(district []DistrictInterface)
	ReplaceDistricts(districts []DistrictInterface)
	AddPoints(points int)
	HasCrown() bool
}

type Player struct {
	ID                string              `json:"id"`
	BuiltDistricts    []DistrictInterface `json:"built_districts"`
	CurrentHand       []DistrictInterface `json:"current_hand"`
	Gold              int                 `json:"gold"`
	SelectedCharacter CharacterInterface  `json:"selected_character"`
	Points            int                 `json:"points"`
	Crown             atomic.Bool         `json:"-"`
	conn              *websocket.Conn     `json:"-"`
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

func (p *Player) GetCurrentHand() []DistrictInterface {
	return p.CurrentHand
}

func (p *Player) SetSelectedCharacter(c CharacterInterface) {
	p.SelectedCharacter = c
}

func (p *Player) GetBuiltDistricts() []DistrictInterface {
	return p.BuiltDistricts
}

// AddDistrict adds a new district to the player
func (p *Player) BuildDistrict(district DistrictInterface) {
	// TODO: Check if the player was able to buy the district
	p.BuiltDistricts = append(p.BuiltDistricts, district)
}

// ReplaceDistricts replaces all districts of the player with new ones
func (p *Player) ReplaceDistricts(districts []DistrictInterface) {
	p.BuiltDistricts = districts
	for _, d := range p.BuiltDistricts {
		d.SetPlayer(p)
	}
}

func (p *Player) AddPoints(points int) {
	p.Points += points
}

func (p *Player) GetPoints() int {
	return p.Points
}

func (p *Player) HasCrown() bool {
	return p.Crown.Load()
}

func (p *Player) SetCrown(has bool) {
	p.Crown.Store(has)
}

func (p *Player) DrawDistricts(d []DistrictInterface) {
	p.CurrentHand = append(p.CurrentHand, d...)
}

// TODO: Temporary location for player creation
func NewPlayer(conn *websocket.Conn) PlayerInterface {
	return &Player{
		ID:                uuid.New().String(),
		BuiltDistricts:    make([]DistrictInterface, 20),
		CurrentHand:       []DistrictInterface{},
		Gold:              2,
		SelectedCharacter: nil,
		Points:            0,
		conn:              conn,
	}
}
