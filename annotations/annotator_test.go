package annotations

import (
	"testing"
)

func setupAnnotatorTest() []Annotation {
	return []Annotation{
		{Tag: LINE},
		{Tag: BUFFER},
		{Tag: WORD},
		{Tag: LINE},
	}
}

func TestOfTag(t *testing.T) {
	setup := setupAnnotatorTest()
	result := OfTag(setup, LINE)
	if len(result) != 2 {
		t.Errorf("Expect 2 results")
	}
}
