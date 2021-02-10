package annotations_test

import (
	"testing"

	"github.com/axiomabsolute/gramme/annotations"
)

func TestGetAnnotatedText(t *testing.T) {
	testText := "These are words."
	wordAnnotations := annotations.AnnotateWords(testText)
	expect := []string{" ", "are", " "}

	actual := annotations.GetAnnotatedText(testText, wordAnnotations[1])

	for i, component := range actual {
		if component != expect[i] {
			t.Errorf("Annotated text should match: Got `%s`, expected `%s`.", component, expect)
		}
	}
}
