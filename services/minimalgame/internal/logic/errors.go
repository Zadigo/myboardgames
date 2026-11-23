package logic

type GameErrors string

func (e GameErrors) Error() string {
	return string(e)
}

const (
	PlayerAlreadyExistsError GameErrors = "player exists"
	PlayerNotFoundError      GameErrors = "player not found"
)
