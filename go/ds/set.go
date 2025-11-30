package ds

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

func (s Set[T]) Contains(element T) bool {
	_, ok := s[element]
	return ok
}
