package main

import (
	"errors"
	"fmt"
)

// Generic Queue Interface
type Queuer[T any] interface {
	Enqueue(T)
	Dequeue() (T, error)
	IsEmpty() bool
}

// Queue structure
type queue[T any] struct {
	data []T
}

// Constructor
func NewQueue[T any]() Queuer[T] {
	return &queue[T]{
		data: make([]T, 0),
	}
}

// Add item to queue (FIFO)
func (q *queue[T]) Enqueue(item T) {
	q.data = append(q.data, item)
}

// Remove first item
func (q *queue[T]) Dequeue() (T, error) {
	var zero T

	if q.IsEmpty() {
		return zero, errors.New("queue is empty")
	}

	item := q.data[0]
	q.data = q.data[1:]

	return item, nil
}

// Check if empty
func (q *queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func main() {
	q := NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue()) // error
}
