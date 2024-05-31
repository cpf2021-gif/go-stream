package stream

import (
	"sort"
)

/*
	Intermediate Operations
*/

func (s stream[E, V]) Filter(predicate func(int, E) bool) stream[E, V] {
	return stream[E, V]{
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

// E -> V
func (s stream[E, V]) Map(mapper func(int, E) V) stream[V, E] {
	return stream[V, E]{
		iter: func(yield func(int, V) bool) {
			for i, v := range s.iter {
				if !yield(i, mapper(i, v)) {
					return
				}
			}
		},
	}
}

func (s stream[E, V]) Sorted(lessfunc func(a, b any) bool) stream[E, V] {
	slice := s.ToSlice()
	sort.Slice(slice, func(i, j int) bool {
		return lessfunc(slice[i], slice[j])
	})
	return StreamOf[[]E, E, V](slice)
}

func (s stream[E, V]) Distinct(fn func(v any) any) stream[E, V] {
	return stream[E, V]{
		iter: func(yield func(int, E) bool) {
			set := make(map[any]struct{})

			for i, v := range s.iter {
				if _, loaded := set[fn(v)]; !loaded {
					set[fn(v)] = struct{}{}
					if !yield(i, v) {
						return
					}
				}
			}
		},
	}
}

func (s stream[E, V]) Head(n uint64) stream[E, V] {

	return stream[E, V]{
		iter: func(yield func(int, E) bool) {
			i := 0
			for _, v := range s.iter {
				if uint64(i) >= n {
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

func (s stream[E, V]) BackWalk() stream[E, V] {
	return stream[E, V]{
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

func (s stream[E, V]) Tail(n uint64) stream[E, V] {

	return stream[E, V]{
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
