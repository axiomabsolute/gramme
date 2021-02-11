package keymap

import "github.com/axiomabsolute/gramme/comlang/pos"

// KeyMap - Map from PoS -> key binding -> word name
type KeyMap map[pos.Tag]map[string]string
