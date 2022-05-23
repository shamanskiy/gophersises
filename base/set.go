package base

import "sync"

type Set[T comparable] map[T]struct{}

func MakeSet[T comparable]() Set[T] {
	return map[T]struct{}{}
}

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

type ThreadSafeSet[T comparable] struct {
	set  Set[T]
	lock sync.RWMutex
}

func MakeThreadSafeSet[T comparable]() *ThreadSafeSet[T] {
	return &ThreadSafeSet[T]{set: MakeSet[T]()}
}

func (tsSet *ThreadSafeSet[T]) Add(elem T) {
	tsSet.lock.Lock()
	defer tsSet.lock.Unlock()
	tsSet.set.Add(elem)
}

func (tsSet *ThreadSafeSet[T]) Empty() bool {
	tsSet.lock.RLock()
	defer tsSet.lock.RUnlock()
	return tsSet.set.Empty()
}

func (tsSet *ThreadSafeSet[T]) Has(elem T) bool {
	tsSet.lock.RLock()
	defer tsSet.lock.RUnlock()

	return tsSet.set.Has(elem)
}

func (tsSet *ThreadSafeSet[T]) UnsafeHas(elem T) bool {
	return tsSet.set.Has(elem)
}

func (tsSet *ThreadSafeSet[T]) Remove(elem T) {
	tsSet.lock.Lock()
	defer tsSet.lock.Unlock()

	delete(tsSet.set, elem)
}

func (tsSet *ThreadSafeSet[T]) Next() T {
	tsSet.lock.RLock()
	defer tsSet.lock.RUnlock()

	for elem := range tsSet.set {
		return elem
	}
	var emptyResult T
	return emptyResult
}

func (tsSet *ThreadSafeSet[T]) Pop() T {
	tsSet.lock.Lock()
	defer tsSet.lock.Unlock()

	for elem := range tsSet.set {
		tsSet.set.Remove(elem)
		return elem
	}
	var emptyResult T
	return emptyResult
}

func (tsSet *ThreadSafeSet[T]) UnsafeToSlice() []T {
	return tsSet.set.ToSlice()
}
