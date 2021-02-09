package annotations

import (
	"regexp"

	"github.com/axiomabsolute/gramme/primitives"
)

// AnnotateBuffer - Generates Buffer annotations for text
func AnnotateBuffer(text string) []Annotation {
	textLength := primitives.Cursor(len(text))
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
func AnnotateByDelimiter(tag Tag, delimiter *regexp.Regexp, text string) []Annotation {
	matches := delimiter.FindAllStringIndex(text, -1)
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

// AnnotateLines - Generate annotations for lines, separated by one more more newline characters
func AnnotateLines(text string) []Annotation {
	pattern := regexp.MustCompile("\\A|\n+|\\z")
	return AnnotateByDelimiter(LINE, pattern, text)
}

// AnnotateWords - Generate annotations for wods, separated by punctuation or whitespace
func AnnotateWords(text string) []Annotation {
	pattern := regexp.MustCompile("\\A|\\n|[\\.,!?]?\\s+|\\z")
	return AnnotateByDelimiter(WORD, pattern, text)
}

// BatchAnnotate - Given a set of annotators, run each on the text and merge results
func BatchAnnotate(annotators []Annotator, text string) []Annotation {
	results := []Annotation{}
	for _, annotator := range annotators {
		annotations := annotator(text)
		results = append(results, annotations...)
	}
	return results
}
