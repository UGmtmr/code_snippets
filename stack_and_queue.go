package codesnippets

type Stack[T any] []T

func (s *Stack[T]) Pop() T {
	n := len(*s)
	v := (*s)[n-1]
	*s = (*s)[:n-1]
	return v
}

type Queue[T any] []T

func (q *Queue[T]) Pop() T {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}
