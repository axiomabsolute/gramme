package annotations

import (
	"regexp"

	"github.com/axiomabsolute/gramme/primitives"
)

// Batch - An Annotator which annotates an entire text greedily
type Batch struct {
	text        *string
	annotations []Annotation
}

// NewBatch - Create a Batch and run annotators
func NewBatch(text *string) *Batch {
	batch := Batch{text: text}
	batch.BatchAnnotate([]AnnotationRule{annotateBuffer, annotateLines, annotateWords})
	return &batch
}

func (b Batch) Containing(cursor primitives.Cursor) []Annotation {
	results := []Annotation{}
	for _, annotation := range b.annotations {
		if annotation.Left.A <= cursor && annotation.Right.B >= cursor {
			results = append(results, annotation)
		}
	}
	return results
}

// All - Returns all annotations
func (b Batch) All() []Annotation {
	return b.annotations
}

// AnnotateBuffer - Generates Buffer annotations for text
func annotateBuffer(text *string) []Annotation {
	textLength := primitives.Cursor(len(*text))
	return []Annotation{
		{
			Tag: BUFFER,
			Region: primitives.Region{
				Left:  primitives.Span{A: 0, B: 0},
				Right: primitives.Span{A: textLength, B: textLength},
			},
		},
	}
}

// AnnotateByDelimiter - Given a pattern that defines a delimiter, annotate the given text into overlapping regions
func annotateByDelimiter(tag Tag, delimiter *regexp.Regexp, text *string) []Annotation {
	matches := delimiter.FindAllStringIndex(*text, -1)
	results := []Annotation{}
	for i := 0; i < len(matches)-1; i++ {
		left := matches[i]
		right := matches[i+1]

		leftSpan := primitives.Span{A: primitives.Cursor(left[0]), B: primitives.Cursor(left[1])}
		rightSpan := primitives.Span{A: primitives.Cursor(right[0]), B: primitives.Cursor(right[1])}

		annotation := Annotation{Tag: tag, Region: primitives.Region{Left: leftSpan, Right: rightSpan}}
		results = append(results, annotation)
	}
	return results
}

// annotateLines - Generate annotations for lines, separated by one more more newline characters
func annotateLines(text *string) []Annotation {
	pattern := regexp.MustCompile("\\A|\n+|\\z")
	return annotateByDelimiter(LINE, pattern, text)
}

// annotateWords - Generate annotations for wods, separated by punctuation or whitespace
func annotateWords(text *string) []Annotation {
	pattern := regexp.MustCompile("\\A|\\n|[\\.,!?]?(?:\\s+|\\z)")
	return annotateByDelimiter(WORD, pattern, text)
}

// BatchAnnotate - Given a set of annotators, run each on the text and merge results
func (b *Batch) BatchAnnotate(annotators []AnnotationRule) []Annotation {
	results := []Annotation{}
	for _, annotator := range annotators {
		annotations := annotator(b.text)
		results = append(results, annotations...)
	}
	b.annotations = results
	return results
}
