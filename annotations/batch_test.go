package annotations

import (
	"testing"
)

func setupBatchTest() *Batch {
	b := NewBatch(&testText)
	return b
}

func TestAll(t *testing.T) {
	b := setupBatchTest()
	all := b.All()
	expectedCount := len(testAnnotatedTexts)
	if len(all) != expectedCount {
		t.Errorf("Expected BatchAnnotator.All to return %d result, got %d", expectedCount, len(all))
	}

	uniqueAnnotations := GetTestAnnotationsMap()
	for _, annotation := range all {
		if !uniqueAnnotations[annotation] {
			t.Errorf("Extracted annotation %v not in expected annotations", annotation)
		}
	}

}
