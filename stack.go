package main

import "math"

type snode[T any] struct {
	value T
	prev  *snode[T]
}

type stack[T any] struct {
	length int
	head   *snode[T]
}

func newStack[T any]() *stack[T] {
	return &stack[T]{
		length: 0,
		head:   nil,
	}
}

func (s *stack[T]) push(item T) {
	n := &snode[T]{value: item}
	s.length++
	if s.head == nil {
		s.head = n
		return
	}

	n.prev = s.head
	s.head = n
}

func (s *stack[T]) pop() T {
	s.length = int(math.Max(0.0, float64(s.length-1)))
	if s.length == 0 {
		var result T
		head := s.head
		s.head = nil
		if head != nil {
			result = head.value
		}
		return result
	}

	head := s.head
	s.head = head.prev
	return head.value
}

func (s *stack[T]) peek() T {
	var result T
	if s.head != nil {
		result = s.head.value
	}
	return result
}
