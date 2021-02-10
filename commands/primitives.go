package commands

// PartOfSpeechTag - Distinct types of parts of speech
//go:generate enumer -type=PartOfSpeechTag -json -text
type PartOfSpeechTag int

const (
	Verb PartOfSpeechTag = iota
	Noun
	TextObject
	Preposition
)

type Grammar struct {
	Name string
}
