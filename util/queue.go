package util

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Add(e ...T) {
	q.elements = append(q.elements, e...)
}

func (q *Queue[T]) Poll() T {
	elem := q.elements[0]
	q.elements = q.elements[1:]
	return elem
}

func (q *Queue[T]) Elements() []T {
	return q.elements
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}
