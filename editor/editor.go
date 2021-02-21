package editor

import (
	"github.com/axiomabsolute/gramme/annotator"
	"github.com/axiomabsolute/gramme/annotator/annotation"
	"github.com/axiomabsolute/gramme/primitives"
)

// Mode - Enum for editor modes
type Mode int

// Editor modes
//go:generate enumer -type=Mode -json -text
const (
	Unkonwn Mode = iota
	ErrorMode
	NormalMode
	InsertMode
)

func standardRegisterKeys() []string {
	return []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "*", "+",
	}
}

func standardRegisters() map[string][]string {
	result := make(map[string][]string)
	for _, v := range standardRegisterKeys() {
		result[v] = make([]string, 0)
	}
	return result
}

// Editor - Model for editor state
type Editor struct {
	text        string
	selections  []primitives.Selection
	registers   map[string][]string
	mode        Mode
	annotators  []annotator.Annotator
	annotations []annotation.Annotation
}

func NewEditor(text string) Editor {
	annotators := annotator.StandardAnnotators()
	annotations := annotator.BatchAnnotate(annotators, text)
	return Editor{
		text:        text,
		selections:  make([]primitives.Selection, 0),
		registers:   make(map[string][]string),
		mode:        NormalMode,
		annotators:  annotators,
		annotations: annotations,
	}
}

func (e Editor) QueryAnnotation(c primitives.Cursor, t annotation.Tag) annotation.Annotation {
	containing := annotation.Containing(e.annotations, c)
	for _, v := range containing {
		if v.Tag == t {
			return v
		}
	}
}

func (e Editor) QueryAnnotations(c primitives.Cursor) []annotation.Annotation {
	containing := annotation.Containing(e.annotations, c)
	results := make([]annotation.Annotation, 0)
	for _, v := range containing {
		results = append(results, v)
	}
	return results
}
