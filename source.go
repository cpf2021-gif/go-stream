package stream

/*
	Source Operations
*/

func StreamOf[Slicens ~[]E, E, V any](source Slicens) stream[E, V] {
	return stream[E, V]{
		iter: func(yield func(int, E) bool) {
			for i, v := range source {
				if !yield(i, v) {
					return
				}
			}
		},
	}
}

func Chunk[Slice ~[]E, E, V any](source Slice, n uint) stream[Slice, V] {
	return stream[Slice, V]{
		iter: func(yield func(int, Slice) bool) {
			if n == 0 {
				return
			}

			num := 0

			for i := uint(0); i < uint(len(source)); i += n {
				end := min(n, uint(len(source[i:])))

				if !yield(num, source[i:i+end:i+end]) {
					return
				}

				num++
			}
		},
	}
}

func GroupBy[Slice ~[]E, E, V any](source Slice, key func(i int, v E) any) stream[Slice, V] {
	groups := make(map[any]Slice)
	for i, v := range source {
		groups[key(i, v)] = append(groups[key(i, v)], v)
	}

	return stream[Slice, V]{
		iter: func(yield func(int, Slice) bool) {
			i := 0
			for _, v := range groups {
				if !yield(i, v) {
					return
				}
				i++
			}
		},
	}
}

// Stream[E, V] -> Stream[T, E]   E -> T
func Map[E, V, T any](source stream[E, V], mapper func(int, E) T) stream[T, E] {
	return stream[T, E]{
		iter: func(yield func(int, T) bool) {
			for i, v := range source.iter {
				if !yield(i, mapper(i, v)) {
					return
				}
			}
		},
	}
}
