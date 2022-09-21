package main

import "testing"

func TestStack(t *testing.T) {
	t.Parallel()

	list := newStack[int]()

	list.push(5)
	list.push(7)
	list.push(9)

	if v := list.pop(); v != 9 {
		t.Errorf("got %v, want 9", v)
	}
	if list.length != 2 {
		t.Errorf("got %v, want 2", list.length)
	}

	list.push(11)
	if v := list.pop(); v != 11 {
		t.Errorf("got %v, want 11", v)
	}
	if v := list.pop(); v != 7 {
		t.Errorf("got %v, want 7", v)
	}
	if v := list.peek(); v != 5 {
		t.Errorf("got %v, want 5", v)
	}
	if v := list.pop(); v != 5 {
		t.Errorf("got %v, want 5", v)
	}
	if v := list.pop(); v != 0 {
		t.Errorf("got %v, want 0", v)
	}

	list.push(69)
	if v := list.peek(); v != 69 {
		t.Errorf("got %v, want 0", v)
	}
	if list.length != 1 {
		t.Errorf("got %v, want 11", list.length)
	}
}
