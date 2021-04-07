package annotations

import (
	"encoding/json"
	"testing"

	"github.com/axiomabsolute/gramme/primitives"
)

var testText string = `The limerick packs laughs anatomical
Into space that is quite economical.
But the good ones I've seen
So seldom are clean
And the clean ones so seldom are comical.`

var annotatedTextsMap map[[3]string]bool
var annotatedTexts [][3]string = [][3]string{
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

var testAnnotations map[Annotation]bool
var testAnnotationsMarshalled []string = []string{
	`{"Tag":"BUFFER","Left":{"A":0,"B":0},"Right":{"A":163,"B":163}}`,
	`{"Tag":"LINE","Left":{"A":0,"B":0},"Right":{"A":36,"B":37}}`,
	`{"Tag":"LINE","Left":{"A":36,"B":37},"Right":{"A":73,"B":74}}`,
	`{"Tag":"LINE","Left":{"A":73,"B":74},"Right":{"A":101,"B":102}}`,
	`{"Tag":"LINE","Left":{"A":101,"B":102},"Right":{"A":121,"B":122}}`,
	`{"Tag":"LINE","Left":{"A":121,"B":122},"Right":{"A":163,"B":163}}`,
	`{"Tag":"WORD","Left":{"A":0,"B":0},"Right":{"A":3,"B":4}}`,
	`{"Tag":"WORD","Left":{"A":3,"B":4},"Right":{"A":12,"B":13}}`,
	`{"Tag":"WORD","Left":{"A":12,"B":13},"Right":{"A":18,"B":19}}`,
	`{"Tag":"WORD","Left":{"A":18,"B":19},"Right":{"A":25,"B":26}}`,
	`{"Tag":"WORD","Left":{"A":25,"B":26},"Right":{"A":36,"B":37}}`,
	`{"Tag":"WORD","Left":{"A":36,"B":37},"Right":{"A":41,"B":42}}`,
	`{"Tag":"WORD","Left":{"A":41,"B":42},"Right":{"A":47,"B":48}}`,
	`{"Tag":"WORD","Left":{"A":47,"B":48},"Right":{"A":52,"B":53}}`,
	`{"Tag":"WORD","Left":{"A":52,"B":53},"Right":{"A":55,"B":56}}`,
	`{"Tag":"WORD","Left":{"A":55,"B":56},"Right":{"A":61,"B":62}}`,
	`{"Tag":"WORD","Left":{"A":61,"B":62},"Right":{"A":72,"B":74}}`,
	`{"Tag":"WORD","Left":{"A":72,"B":74},"Right":{"A":77,"B":78}}`,
	`{"Tag":"WORD","Left":{"A":77,"B":78},"Right":{"A":81,"B":82}}`,
	`{"Tag":"WORD","Left":{"A":81,"B":82},"Right":{"A":86,"B":87}}`,
	`{"Tag":"WORD","Left":{"A":86,"B":87},"Right":{"A":91,"B":92}}`,
	`{"Tag":"WORD","Left":{"A":91,"B":92},"Right":{"A":96,"B":97}}`,
	`{"Tag":"WORD","Left":{"A":96,"B":97},"Right":{"A":101,"B":102}}`,
	`{"Tag":"WORD","Left":{"A":101,"B":102},"Right":{"A":104,"B":105}}`,
	`{"Tag":"WORD","Left":{"A":104,"B":105},"Right":{"A":111,"B":112}}`,
	`{"Tag":"WORD","Left":{"A":111,"B":112},"Right":{"A":115,"B":116}}`,
	`{"Tag":"WORD","Left":{"A":115,"B":116},"Right":{"A":121,"B":122}}`,
	`{"Tag":"WORD","Left":{"A":121,"B":122},"Right":{"A":125,"B":126}}`,
	`{"Tag":"WORD","Left":{"A":125,"B":126},"Right":{"A":129,"B":130}}`,
	`{"Tag":"WORD","Left":{"A":129,"B":130},"Right":{"A":135,"B":136}}`,
	`{"Tag":"WORD","Left":{"A":135,"B":136},"Right":{"A":140,"B":141}}`,
	`{"Tag":"WORD","Left":{"A":140,"B":141},"Right":{"A":143,"B":144}}`,
	`{"Tag":"WORD","Left":{"A":143,"B":144},"Right":{"A":150,"B":151}}`,
	`{"Tag":"WORD","Left":{"A":150,"B":151},"Right":{"A":154,"B":155}}`,
	`{"Tag":"WORD","Left":{"A":154,"B":155},"Right":{"A":162,"B":163}}`,
}

func GetTestAnnotatedTextMap() map[[3]string]bool {
	if annotatedTextsMap != nil {
		return annotatedTextsMap
	}
	result := map[[3]string]bool{}
	for _, v := range annotatedTexts {
		result[v] = true
	}
	annotatedTextsMap = result
	return result
}

func GetTestAnnotationsMap() map[Annotation]bool {
	if testAnnotations != nil {
		return testAnnotations
	}
	results := map[Annotation]bool{}
	for _, m := range testAnnotationsMarshalled {
		annotation := Annotation{}
		json.Unmarshal([]byte(m), &annotation)
		results[annotation] = true
	}
	testAnnotations = results
	return results
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
