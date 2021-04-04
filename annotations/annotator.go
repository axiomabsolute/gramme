package annotations

import "github.com/axiomabsolute/gramme/primitives"

// Annotator - A collection of AnnotationRule definitions and a strategy for applying them to a given
// text. Each implementation is responsible for defining how and when the rules are applied.
// Implementations are expected to annotate the following features:
//     - Buffer: The full text
//     - Line: Newline-separated strings
//     - Word: Sequences of text separated by whitespace, newlines, start of file, end of file, or any of ,!?.
type Annotator interface {
	Containing(primitives.Cursor) []Annotation
}

// OfTag - Filters a slice of annotations by Tag
func OfTag(annotations []Annotation, tag Tag) []Annotation {
	results := []Annotation{}
	for _, annotation := range annotations {
		if annotation.Tag == tag {
			results = append(results, annotation)
		}
	}
	return results
}
