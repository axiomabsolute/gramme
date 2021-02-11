package editor

import "github.com/axiomabsolute/gramme/primitives"

// Mode - Enum for editor modes
type Mode int

// Editor modes
//go:generate enumer -type=Mode -json -text
const (
	ErrorMode Mode = iota
	NormalMode
	InsertMode
)

// Editor - Model for editor state
type Editor struct {
	text       string
	selections []primitives.Selection
	registers  map[string][]string
	mode       Mode
}
