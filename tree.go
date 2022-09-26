package main

type binaryNode[T comparable] struct {
	value T
	left, right *binaryNode[T]
}
