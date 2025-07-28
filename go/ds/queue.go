package ds

type Queue[T any] []T

func (q *Queue[T]) Push(elem T) {
	*q = append(*q, elem)
}

func (q *Queue[T]) PushSlice(s []T) {
	*q = append(*q, s...)
}

func (q *Queue[T]) Pop() T {
	if len(*q) == 0 {
		panic("empty queue")
	}
	popped := (*q)[0]
	*q = (*q)[1:]
	return popped
}

func (q *Queue[T]) Len() int {
	return len(*q)
}
