package models

type DistrictInterface interface {
	GetId() string
	GetPlayer() PlayerInterface
	SetPlayer(player PlayerInterface)
	Buy(player PlayerInterface) error
}

type BaseDistrict struct {
	ID     string
	Name   string
	Cost   int
	Color  string
	Level  int
	Symbol string
	Player PlayerInterface
}
