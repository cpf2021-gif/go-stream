package stream

/*
	Terminal Operations
*/

func (s stream[E, V]) ForEach(action func(int, E)) {
	for i, v := range s.iter {
		action(i, v)
	}
}

func (s stream[E, V]) Reduce(reducer func(acc, cur E) E, empty E) E {
	var result E = empty
	for _, v := range s.iter {
		result = reducer(result, v)
	}
	return result
}

func (s stream[E, V]) ToSlice() []E {
	var result []E
	s.ForEach(func(_ int, v E) {
		result = append(result, v)
	})
	return result
}

func (s stream[E, V]) Count() int {
	var count int
	for _, _ = range s.iter {
		count++
	}
	return count
}

func (s stream[E, V]) First() any {
	for _, v := range s.iter {
		return v
	}
	return nil
}

func (s stream[E, V]) Last() any {
	for _, v := range s.iter {
		return v
	}
	return nil
}

func (s stream[E, V]) AllMatch(predicate func(int, E) bool) bool {
	for i, v := range s.iter {
		if !predicate(i, v) {
			return false
		}
	}
	return true
}

func (s stream[E, V]) AnyMatch(predicate func(int, E) bool) bool {
	for i, v := range s.iter {
		if predicate(i, v) {
			return true
		}
	}
	return false
}

func (s stream[E, V]) Max(comparator func(a, b E) bool, empty E) E {
	var max E = empty
	for _, v := range s.iter {
		if comparator(v, max) {
			max = v
		}
	}
	return max
}

func (s stream[E, V]) Min(comparator func(a, b E) bool, empty E) E {
	var min E = empty
	for _, v := range s.iter {
		if comparator(v, min) {
			min = v
		}
	}
	return min
}
