package lexicon

import "github.com/axiomabsolute/gramme/comlang/pos"

// Lexicon - Map from PoS -> word name -> implementation
type Lexicon map[pos.Tag]map[string]interface{}

// DefaultLexicon - Construct the default lexicon
func DefaultLexicon() Lexicon {
	return map[pos.Tag]map[string]interface{}{
		pos.Verb: {
			"change": "",
			"delete": "",
			"exit":   "",
		},
		pos.Preposition: {
			"inside": "",
			"around": "",
		},
		pos.TextObject: {
			"buffer": "",
			"line":   "",
			"word":   "",
		},
	}
}
