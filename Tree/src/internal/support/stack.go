package support


type Stack[T any] struct {
	Arr []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(element T) {
	s.Arr = append(s.Arr, element)
}

func (s *Stack[T]) Pop() (T) {
	if len(s.Arr) == 0 {
		var zero T
		return zero
	}
	topIndex := len(s.Arr) - 1
	element := s.Arr[topIndex]
	s.Arr = s.Arr[:topIndex]
	return element
}
