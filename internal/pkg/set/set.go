package set

import "iter"

type empty struct{}

type Set[T comparable] struct {
	m map[T]empty
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]empty),
	}
}

func (s *Set[T]) Has(value T) bool {
	_, ok := s.m[value]
	return ok
}

func (s *Set[T]) Add(value T) bool {
	if s.Has(value) {
		return false
	}
	s.m[value] = empty{}
	return true
}

func (s *Set[T]) Remove(value T) bool {
	if !s.Has(value) {
		return false
	}
	delete(s.m, value)
	return true
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}
