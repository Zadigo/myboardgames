package games

import (
	"errors"
	"fmt"
)

// StandardGameError is a custom error type for handling errors in a standardized way
// for the game mini servers and its game instances.
type StandardGameError struct {
}

func (e *StandardGameError) BasicError(errs ...error) error {
	return errors.Join(errs...)
}

func (e *StandardGameError) GameNotFound(id string) error {
	return e.BasicError(fmt.Errorf("Game with id %s was not found", id))
}
