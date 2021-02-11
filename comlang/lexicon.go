package lexicon

import "github.com/axiomabsolute/gramme/comlang/pos"

// Lexicon - Map from PoS -> word name -> implementation
type Lexicon map[pos.Tag]map[string]interface{}

func DefaultLexicon() Lexicon {
	return map[pos.Tag]map[string]interface{}{
		pos.Verb: map[string]interface{}{
			"change": "",
			"delete": "",
			"exit":   "",
		},
		pos.Preposition: map[string]interface{}{
			"inside": "",
			"around": "",
		},
		pos.TextObject: map[string]interface{}{
			"buffer": "",
			"line":   "",
			"word":   "",
		},
	}
}
