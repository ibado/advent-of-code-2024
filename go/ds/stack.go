package ds

type Stack[T any] []T

func (s *Stack[T]) Push(elem T) {
	*s = append(*s, elem)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		panic("empty stack")
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}

func (s *Stack[T]) Len() int {
	return len(*s)
}
