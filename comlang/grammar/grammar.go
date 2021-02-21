package grammar

import (
	"github.com/axiomabsolute/gramme/comlang/pos"
	"github.com/axiomabsolute/gramme/editor"
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

type Command func(editor.Editor) editor.Editor

type Impl func(interface{}, []interface{}) Command

// Grammar - Defines how verbs combine with complement to form a command
type Grammar struct {
	Complement []pos.Tag
	Impl       Impl
}
