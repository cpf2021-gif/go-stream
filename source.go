package stream

/*
	Source Operations
*/

func StreamOf[Slicens ~[]E, E any](s Slicens) stream[E] {
	return stream[E]{
		iter: func(yield func(int, E) bool) {
			for i, v := range s {
				if !yield(i, v) {
					return
				}
			}
		},
	}
}

func Chunk[Slice ~[]E, E any](source stream[E], n uint) stream[Slice] {
	s := source.ToSlice()

	return stream[Slice]{
		iter: func(yield func(int, Slice) bool) {
			if n == 0 {
				return
			}

			num := 0

			for i := uint(0); i < uint(len(s)); i += n {
				end := min(n, uint(len(s[i:])))

				if !yield(num, s[i:i+end:i+end]) {
					return
				}

				num++
			}
		},
	}
}
