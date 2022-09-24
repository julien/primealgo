package main

import "errors"

type lnode[T comparable] struct {
	value      T
	prev, next *lnode[T]
}

var insertionError = errors.New("can't insert at specified index")

type LinkedList[T comparable] struct {
	length     int
	head, tail *lnode[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{length: 0, head: nil, tail: nil}
}

func (l LinkedList[T]) Length() int {
	return l.length
}

func (l *LinkedList[T]) InsertAt(item T, idx int) error {
	if idx > l.length {
		return insertionError
	}

	if idx == l.length {
		l.Append(item)
		return nil
	} else if idx == 0 {
		l.Prepend(item)
		return nil
	}

	l.length++
	curr := l.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}

	node := &lnode[T]{value: item}
	node.next = curr
	prev := curr.prev
	node.prev = prev
	curr.prev = node
	prev.next = node

	return nil
}

func (l *LinkedList[T]) Remove(item T) (result T) {
	curr := l.head
	for i := 0; curr != nil && i < l.length; i++ {
		if curr.value == item {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return
	}

	l.length--

	if l.length == 0 {
		out := l.head.value
		l.head = nil
		l.tail = nil
		return out
	}

	if curr.prev != nil {
		curr.prev.next = curr.next
	}
	if curr.next != nil {
		curr.next.prev = curr.prev
	}
	if curr == l.head {
		l.head = curr.next
	}
	if curr == l.tail {
		l.tail = curr.prev
	}

	curr.prev = nil
	curr.next = nil

	result = curr.value
	return result
}

func (l *LinkedList[T]) RemoveAt(idx int) T {
	var result T
	curr := l.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}

	if curr == nil {
		return result
	}

	l.length--

	if l.length == 0 {
		out := l.head.value
		l.head = nil
		l.tail = nil
		return out
	}

	if curr.prev != nil {
		curr.prev.next = curr.next
	}
	if curr.next != nil {
		curr.next.prev = curr.prev
	}
	if curr == l.head {
		l.head = curr.next
	}
	if curr == l.tail {
		l.tail = curr.prev
	}

	curr.prev = nil
	curr.next = nil

	result = curr.value
	return result
}

func (l *LinkedList[T]) Append(item T) {
	l.length++

	node := &lnode[T]{value: item}

	if l.tail == nil {
		l.head = node
		l.tail = node
		return
	}

	node.prev = l.tail
	l.tail.next = node
	l.tail = node
}

func (l *LinkedList[T]) Prepend(item T) {
	node := &lnode[T]{value: item}

	l.length++
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l LinkedList[T]) Get(idx int) T {
	curr := l.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}
	return curr.value
}
