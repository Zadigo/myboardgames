package backend

type BaseTile[T any] struct {
	Uuid   string `json:"uuid"`
	Name   string `json:"name"`
	Points int    `json:"points"`
	// The bonus resources required in order to obtain the tile.
	// For example, a tile might require 3 emerald bonuses
	// from cards to be obtained.
	Resources T                                          `json:"resources"`
	Owner     WebsocketClientInterface[WebsocketMessage] `json:"owner"`
}

type WakandaTile struct {
	BaseTile[MarvelCardResources]
}

type AsgardTile struct {
	BaseTile[MarvelCardResources]
}

type AtlantisTile struct {
	BaseTile[MarvelCardResources]
}

type KnowhereTile struct {
	BaseTile[MarvelCardResources]
}

type AttilanTile struct {
	BaseTile[MarvelCardResources]
}

type AvengersTowerTile struct {
	BaseTile[MarvelCardResources]
}

type TriskelionTile struct {
	BaseTile[MarvelCardResources]
}

type HellsKitchenTile struct {
	BaseTile[MarvelCardResources]
}

type PointsResource struct {
	Points int `json:"points"`
}

type SanctumSanctorumTile struct {
	BaseTile[PointsResource]
}

// Infinity Gauntlet tile is obtained by the first player
// to have 1 resourcce of each type from the Marvel expansion,
// 16 points from bonuses on cards and tiles and a time token bonus
// obtained from cards. The player who obtains the Infinity Gauntlet tile
// wins the game immediately.
type InfinityGauntletTile struct {
	CardResources `json:"resources"`
	TimeToken int `json:"timeToken"`
	Points    int `json:"points"`
}
