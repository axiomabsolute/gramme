package textObject

import (
	"github.com/axiomabsolute/gramme/annotator/annotation"
	"github.com/axiomabsolute/gramme/editor"
	"github.com/axiomabsolute/gramme/primitives"
)

type TextObject func(primitives.Cursor, editor.Editor) primitives.Region

func textObjectByAnnotation(c primitives.Cursor, e editor.Editor, t annotation.Tag) primitives.Region {
	annotation := e.QueryAnnotation(c, t)
	return annotation.Region
}

func Word(c primitives.Cursor, e editor.Editor) primitives.Region {
	return textObjectByAnnotation(c, e, annotation.WORD)
}

func Buffer(c primitives.Cursor, e editor.Editor) primitives.Region {
	return textObjectByAnnotation(c, e, annotation.BUFFER)
}

func Line(c primitives.Cursor, e editor.Editor) primitives.Region {
	return textObjectByAnnotation(c, e, annotation.LINE)
}
