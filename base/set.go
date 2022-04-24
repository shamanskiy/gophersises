package base

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

func (s Set[T]) Has(elem T) bool {
	_, ok := s[elem]
	return ok
}

func (s Set[T]) Next() T {
	for elem := range s {
		return elem
	}
	var emptyResult T
	return emptyResult
}

func (s Set[T]) Pop() T {
	for elem := range s {
		s.Remove(elem)
		return elem
	}
	var emptyResult T
	return emptyResult
}

func (s Set[T]) ToSlice() []T {
	slice := []T{}
	for key := range s {
		slice = append(slice, key)
	}
	return slice
}

func (s Set[T]) Empty() bool {
	return len(s) == 0
}
