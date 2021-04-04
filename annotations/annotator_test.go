package annotations

import (
	"testing"

	"github.com/axiomabsolute/gramme/primitives"
)

var testText string = `The limerick packs laughs anatomical
Into space that is quite economical.
But the good ones I've seen
So seldom are clean
And the clean ones so seldom are comical.`

var testAnnotatedTexts [][3]string = [][3]string{
	{"", "The limerick packs laughs anatomical\nInto space that is quite economical.\nBut the good ones I've seen\nSo seldom are clean\nAnd the clean ones so seldom are comical.", ""},
	{"", "The limerick packs laughs anatomical", "\n"},
	{"\n", "Into space that is quite economical.", "\n"},
	{"\n", "But the good ones I've seen", "\n"},
	{"\n", "So seldom are clean", "\n"},
	{"\n", "And the clean ones so seldom are comical.", ""},
	{"", "The", " "},
	{" ", "limerick", " "},
	{" ", "packs", " "},
	{" ", "laughs", " "},
	{" ", "anatomical", "\n"},
	{"\n", "Into", " "},
	{" ", "space", " "},
	{" ", "that", " "},
	{" ", "is", " "},
	{" ", "quite", " "},
	{" ", "economical", ".\n"},
	{".\n", "But", " "},
	{" ", "the", " "},
	{" ", "good", " "},
	{" ", "ones", " "},
	{" ", "I've", " "},
	{" ", "seen", "\n"},
	{"\n", "So", " "},
	{" ", "seldom", " "},
	{" ", "are", " "},
	{" ", "clean", "\n"},
	{"\n", "And", " "},
	{" ", "the", " "},
	{" ", "clean", " "},
	{" ", "ones", " "},
	{" ", "so", " "},
	{" ", "seldom", " "},
	{" ", "are", " "},
	{" ", "comical", "."},
}

func TestOfTag(t *testing.T) {
	setup := []Annotation{
		{Tag: LINE},
		{Tag: BUFFER},
		{Tag: WORD},
		{Tag: LINE},
	}
	result := OfTag(setup, LINE)
	if len(result) != 2 {
		t.Errorf("Expect 2 results")
	}
}

// TestAnnotators - Test all implementations of the Annotators interface
// for expected functionality
func TestAnnotators(t *testing.T) {
	annotators := map[string]Annotator{
		"batch": NewBatch(&testText),
	}
	expectedStandardTextFeatures := map[string]bool{
		"space":                                true,
		"Into space that is quite economical.": true,
		testText:                               true,
	}
	standardTextFeaturesPosition := 44
	for annotatorName, annotator := range annotators {
		results := annotator.Containing(primitives.Cursor(standardTextFeaturesPosition))
		if len(results) != len(expectedStandardTextFeatures) {
			t.Errorf("Expect %d annotations at position %d using annotator %s", len(expectedStandardTextFeatures), standardTextFeaturesPosition, annotatorName)
		}
		for _, result := range results {
			resultText := GetAnnotatedText(testText, result)
			if !expectedStandardTextFeatures[resultText[1]] {
				t.Errorf("Expected %s to be in results", resultText)
			}
		}
	}
}
