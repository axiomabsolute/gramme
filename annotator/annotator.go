package annotator

import (
	"regexp"

	"github.com/axiomabsolute/gramme/annotator/annotation"
	"github.com/axiomabsolute/gramme/primitives"
)

// Annotator - A function that analyzes a source text and produces a slice of Annotations
type Annotator func(text string) []annotation.Annotation

// AnnotateBuffer - Generates Buffer annotations for text
func AnnotateBuffer(text string) []annotation.Annotation {
	textLength := primitives.Cursor(len(text))
	return []annotation.Annotation{
		{
			Tag: annotation.BUFFER,
			Region: primitives.Region{
				Left:  primitives.Span{A: 0, B: 0},
				Right: primitives.Span{A: textLength, B: textLength},
			},
		},
	}
}

// AnnotateByDelimiter - Given a pattern that defines a delimiter, annotate the given text into overlapping regions
func AnnotateByDelimiter(tag annotation.Tag, delimiter *regexp.Regexp, text string) []annotation.Annotation {
	matches := delimiter.FindAllStringIndex(text, -1)
	results := []annotation.Annotation{}
	for i := 0; i < len(matches)-1; i++ {
		left := matches[i]
		right := matches[i+1]

		leftSpan := primitives.Span{A: primitives.Cursor(left[0]), B: primitives.Cursor(left[1])}
		rightSpan := primitives.Span{A: primitives.Cursor(right[0]), B: primitives.Cursor(right[1])}

		annotation := annotation.Annotation{Tag: tag, Region: primitives.Region{Left: leftSpan, Right: rightSpan}}
		results = append(results, annotation)
	}
	return results
}

// AnnotateLines - Generate annotations for lines, separated by one more more newline characters
func AnnotateLines(text string) []annotation.Annotation {
	pattern := regexp.MustCompile("\\A|\n+|\\z")
	return AnnotateByDelimiter(annotation.LINE, pattern, text)
}

// AnnotateWords - Generate annotations for words, separated by punctuation or whitespace
func AnnotateWords(text string) []annotation.Annotation {
	pattern := regexp.MustCompile("\\A|\\n|[\\.,!?]?\\s+|\\z")
	return AnnotateByDelimiter(annotation.WORD, pattern, text)
}

// BatchAnnotate - Given a set of annotators, run each on the text and merge results
func BatchAnnotate(annotators []Annotator, text string) []annotation.Annotation {
	results := []annotation.Annotation{}
	for _, annotator := range annotators {
		annotations := annotator(text)
		results = append(results, annotations...)
	}
	return results
}

// StandardAnnotators - Return set of standard annotators
func StandardAnnotators() []Annotator {
	return []Annotator{
		AnnotateWords,
		AnnotateLines,
		AnnotateBuffer,
	}
}
