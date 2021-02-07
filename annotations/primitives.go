package annotations

import "github.com/axiomabsolute/gramme/primitives"

// Tag - A type tag for annotations
type Tag int

const (
	// BUFFER - An entire file's content
	BUFFER Tag = iota
	// LINE - A newline delimited piece of text
	LINE
	// WORD - A whitespace or punctuation delimited sequence of text
	WORD
)

// Annotation - A tagged region of text
type Annotation struct {
	Tag Tag
	primitives.Region
}
