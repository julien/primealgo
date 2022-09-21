package main

type qnode[T any] struct {
	value T
	next  *qnode[T]
}

type queue[T any] struct {
	length     int
	head, tail *qnode[T]
}

func newQueue[T any]() *queue[T] {
	return &queue[T]{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}

func (q *queue[T]) enqueue(item T) {
	n := &qnode[T]{value: item}
	q.length++
	if q.tail == nil {
		q.tail = n
		q.head = n
		return
	}

	q.tail.next = n
	q.tail = n
}

func (q *queue[T]) deque() T {
	var result T

	if q.head != nil {
		q.length--
		head := q.head
		q.head = q.head.next

		head.next = nil

		result = head.value
	}

	if q.length == 0 {
		q.tail = nil
	}

	return result
}

func (q *queue[T]) peek() T {
	var result T
	if q.head != nil {
		result = q.head.value
	}
	return result
}
