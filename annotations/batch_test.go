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
	expectedCount := 35
	if len(all) != expectedCount {
		t.Errorf("Expected BatchAnnotator.All to return %d result, got %d", expectedCount, len(all))
	}
}
