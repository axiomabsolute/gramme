package verb

import (
	"github.com/axiomabsolute/gramme/editor"
	"github.com/axiomabsolute/gramme/primitives"
)

func Change(s primitives.Span, e editor.Editor) editor.Editor {
	// Need to delete text and change mode, but also need to handle ADJUSTING other cursors
	return nil
}
