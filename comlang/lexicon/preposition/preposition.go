package preposition

import "github.com/axiomabsolute/gramme/primitives"

type Preposition func(primitives.Region) primitives.Span

func Inside(r primitives.Region) primitives.Span {
	return primitives.Span{A: r.Left.B, B: r.Right.A}
}

func Around(r primitives.Region) primitives.Span {
	return primitives.Span{A: r.Left.A, B: r.Right.B}
}
