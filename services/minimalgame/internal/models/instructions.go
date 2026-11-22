package models

import "context"

type InstructionMetaData struct {
	Description string // A description that can be added
}

// Instruction represents a game move
type Instruction struct {
	Action   string   // The name of the action to be called
	Caller   []string // The player(s) making the move (the move can also be done by the computer)
	On       []string // The player(s) that will be affected by the move (if applicable)
	TableID  string   // A reference to the table on which the move is applied
	CardID   string   // The object or card involved in the move
	MetaData InstructionMetaData
}

func NewInstruction(ctx context.Context, action string, Caller []string) *Instruction {
	return &Instruction{
		Action:   action,
		Caller:   Caller,
		On:       make([]string, 0),
		TableID:  "",
		CardID:   "",
		MetaData: InstructionMetaData{},
	}
}
