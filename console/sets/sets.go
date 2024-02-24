package sets

import "golang.org/x/exp/maps"

// Set ...
type Set[T comparable] map[T]struct{}

// New ...
func New[T comparable](a ...T) Set[T] {
	s := make(Set[T])
	for _, v := range a {
		s[v] = struct{}{}
	}
	return s
}

// Add adds a value to the set.
func (s Set[T]) Add(i T) {
	s[i] = struct{}{}
}

// Contains returns whether the given value is in the set.
func (s Set[T]) Contains(i T) bool {
	_, ok := s[i]
	return ok
}

// Delete deletes a value in the set.
func (s Set[T]) Delete(i T) {
	delete(s, i)
}

// Values returns all the values currently in the set.
func (s Set[T]) Values() []T {
	return maps.Keys(s)
}

// Clear removes all values from the set.
func (s Set[T]) Clear() {
	for i := range s {
		s.Delete(i)
	}
}
