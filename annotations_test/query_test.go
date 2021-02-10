package annotations_test

import (
	"fmt"
	"testing"

	"github.com/axiomabsolute/gramme/annotator"
	"github.com/axiomabsolute/gramme/annotator/annotation"
	"github.com/axiomabsolute/gramme/primitives"
)

const testMathText = `A dozen, a gross, and a score
Plus three times the square root of four
Divided by seven
Plus five times eleven
Is nine squared and not a bit more.`

func TestOfTag(t *testing.T) {

	annotators := annotator.StandardAnnotators()
	extractedAnnotations := annotator.BatchAnnotate(annotators, testMathText)

	expectedLength := 5
	expectedFirstLine := "A dozen, a gross, and a score"

	lineAnnotations := annotation.OfTag(extractedAnnotations, annotation.LINE)

	if len(lineAnnotations) != expectedLength {
		t.Errorf("Should annotate each line")
	}

	if annotation.GetAnnotatedText(testMathText, lineAnnotations[0])[1] != expectedFirstLine {
		t.Errorf("Should be in order and include the first line of text")
	}
}

func TestContaining(t *testing.T) {
	annotators := annotator.StandardAnnotators()
	extractedAnnotations := annotator.BatchAnnotate(annotators, testMathText)

	expectedLength := 3

	lineAnnotations := annotation.Containing(extractedAnnotations, primitives.Cursor(13))

	for _, v := range lineAnnotations {
		text := annotation.GetAnnotatedText(testMathText, v)
		fmt.Printf("%s", text)
	}

	if len(lineAnnotations) != expectedLength {
		t.Errorf("Should extract %d annotations, found %d", expectedLength, len(lineAnnotations))
	}
}
