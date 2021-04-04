package annotations

import "github.com/axiomabsolute/gramme/primitives"

// Annotator - A collection of AnnotationRule definitions and a
// strategy for applying them to a given text. Each implementation is
// responsible for defining how and when the rules are applied.
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
