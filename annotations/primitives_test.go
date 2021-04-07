package annotations

import (
	"testing"

	"github.com/axiomabsolute/gramme/primitives"
)

func TestGetAnnotatedText(t *testing.T) {
	text := "foo bar baz"
	annotation := Annotation{Tag: WORD, Region: primitives.Region{Left: primitives.Span{A: 3, B: 4}, Right: primitives.Span{A: 7, B: 8}}}
	result := GetAnnotatedText(text, annotation)
	expectedInner := "bar"
	if len(result) != 3 {
		t.Errorf("Expected 3 results; prefix inner suffix")
	}
	if result[0] != " " {
		t.Errorf("Expected prefix ` `")
	}
	if result[1] != expectedInner {
		t.Errorf("Expected inner `%s`", expectedInner)
	}
	if result[2] != " " {
		t.Errorf("Expected suffix ` `")
	}
}

func TestAllGetAnnotatedText(t *testing.T) {
	annotationsMap := GetTestAnnotationsMap()
	annotatedTextsMap := GetTestAnnotatedTextMap()
	for annotation, _ := range annotationsMap {
		text := GetAnnotatedText(testText, annotation)
		if !annotatedTextsMap[text] {
			t.Errorf("Expected texts %v", text)
		}
	}
}
