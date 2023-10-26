package ADT

type Stack[T AlNum] struct {
	data []T
}

func (s Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s Stack[T]) Pop() T {
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s Stack[T]) Peek() T {
	return s.data[len(s.data)-1]
}
