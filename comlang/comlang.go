package comlang

import (
	"github.com/axiomabsolute/gramme/comlang/grammar"
	"github.com/axiomabsolute/gramme/comlang/pos"
	"github.com/axiomabsolute/gramme/editor"
)

type UnaryVerb func(editor.Editor) editor.Editor
type PrepTOVerb func() editor.Editor

var GrammarRegistry grammar.Registry = map[grammar.Name]grammar.Grammar{
	grammar.Unary: {
		Complement: []pos.Tag{},
		Impl: func(verb interface{}, complement []interface{}) grammar.Command {
			resolvedVerb := verb.(UnaryVerb)
			return resolvedVerb
		},
	},
	grammar.PrepTextObject: {
		Complements: []pos.Tag{pos.Preposition, pos.TextObject},
		Impl: func(verb interface{}, complement []interface{}) grammar.Command {
			resolvedVerb := verb.(PrepTOVerb)
		}
	}
}
