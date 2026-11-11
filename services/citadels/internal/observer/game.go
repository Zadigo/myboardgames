package observer

type GameSubject struct {
	BaseSubject
}

func (g *GameSubject) Update(message ObserverMessage) error {
	return nil
}
