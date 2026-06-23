package cards

type CardInterface interface {
	GetUuid() string
	// Resolution logic for the card's effect, if any
	Resolve(currentScore int, options ResolveOptions) int
	SetPlayer(playerUuid string)
}

type CardFactory interface {
	createCards() []CardInterface
}
