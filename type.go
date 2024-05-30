package stream

import (
	"iter"
)

type stream[E any] struct {
	iter iter.Seq2[int, E]
}
