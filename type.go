package stream

import (
	"iter"
)

type stream[E, V any] struct {
	iter   iter.Seq2[int, E]
	target V
}
