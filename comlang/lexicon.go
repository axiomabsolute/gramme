package lexicon

import "github.com/axiomabsolute/gramme/comlang/pos"

// Lexicon - Map from PoS -> word name -> implementation
type Lexicon map[pos.Tag]map[string]interface{}
