package annotation

import (
	"github.com/axiomabsolute/gramme/primitives"
)

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

// Containing - Filters a slice of annotations by cursor inclusion
func Containing(annotations []Annotation, cursor primitives.Cursor) []Annotation {
	results := []Annotation{}
	for _, annotation := range annotations {
		if annotation.Left.A <= cursor && annotation.Right.B >= cursor {
			results = append(results, annotation)
		}
	}
	return results
}
