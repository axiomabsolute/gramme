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

func setupBatchTest() *Batch {
	b := NewBatch(&testText)
	return b
}

func TestAnnotateBuffer(t *testing.T) {
	annotator := setupBatchTest()
	result := OfTag(annotator.All(), BUFFER)
	if len(result) != 1 {
		t.Errorf("AnnotateBuffer(..) should return exactly one annotation")
	}
	bufferAnnotation := result[0]
	if bufferAnnotation.Left.A != 0 {
		t.Errorf("Buffer annotation should start at 0")
	}
	if bufferAnnotation.Left.B != 0 {
		t.Errorf("Buffer annotation left region should be empty")
	}
	if bufferAnnotation.Right.A != primitives.Cursor(len(testText)) {
		t.Errorf("Buffer annotation should end at the full text")
	}
	if bufferAnnotation.Right.B != bufferAnnotation.Right.A {
		t.Errorf("Buffer annotation right region should be empty")
	}
}

func TestAnnotateLines(t *testing.T) {
	annotator := setupBatchTest()
	result := OfTag(annotator.All(), LINE)
	if len(result) != 5 {
		t.Errorf("AnnotateLines(..) should return 5 lines")
	}
}

func TestAnnotateWords(t *testing.T) {
	annotator := setupBatchTest()
	result := OfTag(annotator.All(), WORD)
	if len(result) != 29 {
		t.Errorf("AnnotateWords(..) should return 29 words")
	}
	spotCheck := GetAnnotatedText(testText, result[10])
	expect := []string{
		` `,
		`economical`,
		`.
`,
	}
	for i, expected := range expect {
		if spotCheck[i] != expected {
			t.Errorf("AnnotateWords(..)[10] - Expected result[%v] to be `%v` but found `%v`", i, expected, spotCheck[i])
		}
	}

}
