package main

import (
	"errors"
	"fmt"
	//"golang.org/x/exp/constraints"
)

// Generic Stack Interface
type Stacker[T any] interface {
	Push(t T)
	Pop() (T, error)
	IsEmpty() bool
}

// Stack structure
type stack[T any] struct {
	data []T
}

// Constructor
func NewStack[T any]() Stacker[T] {
	return &stack[T]{
		data: make([]T, 0),
	}
}

// Push adds item to stack
func (s *stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

// Pop removes last item
func (s *stack[T]) Pop() (T, error) {
	var zero T
	//var t T

	if s.IsEmpty() {
		return zero, errors.New("the stack is empty")
	}

	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]

	return item, nil
}

// IsEmpty checks if stack is empty
func (s *stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func main() {
	// Create stack for int
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop()) // will return error
}
