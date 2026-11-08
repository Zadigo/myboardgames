package models

import (
	"uuid"

	"github.com/Zadigo/citadels/internal/districts"
)

type PlayerInterface interface {
	GetId() string
	GoldCount() int
	IncreaseGold(by int)
	DecreaseGold(by int)
	GetSelectedCharacter() CharacterInterface
	SetSelectedCharacter(c CharacterInterface)
}

type Player struct {
	ID                string
	Districts         []districts.BaseDistrict
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

// TODO: Temporary location for player creation
func NewPlayer() PlayerInterface {
	return &Player{
		ID:                uuid.New().String(),
		Gold:              0,
		Districts:         []districts.BaseDistrict{},
		SelectedCharacter: nil,
	}
}
