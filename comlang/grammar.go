package grammar

import (
	"github.com/axiomabsolute/gramme/comlang/pos"
)

// Name - Grammar name enum
type Name int

// Unique grammar names
//go:generate enumer -type=Name -json -text
const (
	Unknown Name = iota
	Unary
	PrepTextObject
)

// Registry - Index Grammars by name
type Registry map[Name]Grammar

// Grammar - Defines how verbs combine with complement to form a command
type Grammar struct {
	Name       Name
	Complement []pos.Tag
}
