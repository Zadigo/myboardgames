package models

import "github.com/Zadigo/citadels/internal/districts"

type Player struct {
	ID                string
	Districts         []districts.BaseDistrict
	Gold              int
	SelectedCharacter string
}
