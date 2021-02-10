package pos

// Tag - Distinct types of parts of speech
//go:generate enumer -type=Tag -json -text
type Tag int

// Enum for parts of speech
const (
	Verb Tag = iota
	Noun
	TextObject
	Preposition
)
