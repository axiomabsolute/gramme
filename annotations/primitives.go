package annotations

import "github.com/axiomabsolute/gramme/primitives"

// Annotator - A function that analyzes a source text and produces a slice of Annotations
type Annotator func(text string) []Annotation

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

// GetAnnotatedText - Returns 3 string; the left delimiter, annotated region, and right delimiter
func GetAnnotatedText(text string, annotation Annotation) []string {
	left := text[annotation.Left.A:annotation.Left.B]
	middle := text[annotation.Left.B:annotation.Right.A]
	right := text[annotation.Right.A:annotation.Right.B]
	return []string{
		left, middle, right,
	}
}
