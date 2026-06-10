package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(el T) {
	s.elements = append(s.elements, el)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	existingLen := len(s.elements)
	newLen := existingLen - 1
	elementToRemove := s.elements[newLen]

	//Sol 1: With new capacity
	newElements := s.elements[:newLen]
	s.elements = make([]T, 0, len(newElements))
	s.elements = append(s.elements, newElements...)

	//Sol 2: With existing capacity
	// s.elements = s.elements[:newLen]

	return elementToRemove, true

}

func (s *Stack[T]) IsEmpty() bool {
	length := len(s.elements)

	return length == 0
}

func main() {
	s := Stack[string]{
		elements: []string{},
	}

	n := Stack[int]{
		elements: []int{},
	}

	s.Push("Rohit")
	s.Push("Mohit")
	s.Pop()
	s.Push("Purohit")
	s.Push("Mohit")
	s.Pop()
	fmt.Println(s.elements)
	fmt.Println("Cap", cap(s.elements))
	fmt.Println("Len", len(s.elements))

	n.Push(1)
	n.Push(2)
	n.Pop()
	n.Push(3)
	fmt.Println(n.elements)
}
