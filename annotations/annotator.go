package annotations

import "github.com/axiomabsolute/gramme/primitives"

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
