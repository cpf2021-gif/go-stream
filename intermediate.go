package stream

import (
	"sort"
	"sync"
)

/*
	Intermediate Operations
*/

func (s stream[E]) Filter(predicate func(int, E) bool) stream[E] {
	return stream[E]{
		iter: func(yield func(int, E) bool) {
			for i, v := range s.iter {
				if predicate(i, v) {
					if !yield(i, v) {
						return
					}
				}
			}
		},
	}
}

func (s stream[E]) Map(mapper func(int, E) E) stream[E] {
	return stream[E]{
		iter: func(yield func(int, E) bool) {
			for i, v := range s.iter {
				if !yield(i, mapper(i, v)) {
					return
				}
			}
		},
	}
}

func (s stream[E]) Sorted(lessfunc func(a, b any) bool) stream[E] {
	slice := s.ToSlice()
	sort.Slice(slice, func(i, j int) bool {
		return lessfunc(slice[i], slice[j])
	})
	return StreamOf(slice)
}

func (s stream[E]) Distinct(fn func(v any) any) stream[E] {
	return stream[E]{
		iter: func(yield func(int, E) bool) {
			var m sync.Map

			for i, v := range s.iter {
				if _, loaded := m.LoadOrStore(fn(v), struct{}{}); !loaded {
					if !yield(i, v) {
						return
					}
				}
			}
		},
	}
}

func (s stream[E]) Head(n int64) stream[E] {
	if n < 1 {
		panic("n must be greater than 0")
	}

	return stream[E]{
		iter: func(yield func(int, E) bool) {
			i := 0
			for _, v := range s.iter {
				if int64(i) >= n {
					return
				}
				if !yield(i, v) {
					return
				}
				i++
			}
		},
	}
}

func (s stream[E]) BackWalk() stream[E] {
	return stream[E]{
		iter: func(yield func(int, E) bool) {
			slice := s.ToSlice()
			for i := len(slice) - 1; i >= 0; i-- {
				if !yield(i, slice[i]) {
					return
				}
			}
		},
	}
}

func (s stream[E]) Tail(n int64) stream[E] {
	if n < 1 {
		panic("n must be greater than 0")
	}

	return stream[E]{
		iter: func(yield func(int, E) bool) {
			slice := s.ToSlice()
			start := max(len(slice)-int(n), 0)
			for i := start; i < len(slice); i++ {
				if !yield(i, slice[i]) {
					return
				}
			}
		},
	}
}
