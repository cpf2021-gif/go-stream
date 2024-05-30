package stream

/*
	Terminal Operations
*/

func (s stream[E]) ForEach(action func(int, E)) {
	for i, v := range s.iter {
		action(i, v)
	}
}

func (s stream[E]) ToSlice() []E {
	var result []E
	s.ForEach(func(_ int, v E) {
		result = append(result, v)
	})
	return result
}

func (s stream[E]) Count() int {
	var count int
	for _, _ = range s.iter {
		count++
	}
	return count
}

func (s stream[E]) First() any {
	for _, v := range s.iter {
		return v
	}
	return nil
}

func (s stream[E]) Last() any {
	for _, v := range s.iter {
		return v
	}
	return nil
}
