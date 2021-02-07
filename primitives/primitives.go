package primitives

// Cursor - An offset-based location relative to the start of a file
type Cursor int

// Span - A sequence of text represented by the text between two cursor positions.
type Span struct {
	A Cursor
	B Cursor
}

// Region - A delimited sequence of text represented by two non-overlapping spans (the delimiters) and
// the implicit span between them.
type Region struct {
	Left  Span
	Right Span
}
