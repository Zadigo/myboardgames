package games

type GameErrors string

func (e GameErrors) Error() string {
	return string(e)
}

const ()
