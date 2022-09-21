package main

import "testing"

func TestQueue(t *testing.T) {
	t.Parallel()

	list := newQueue[int]()

	list.enqueue(5)
	list.enqueue(7)
	list.enqueue(9)

	if v := list.deque(); v != 5 {
		t.Errorf("got %v, want 5", v)
	}
	if list.length != 2 {
		t.Errorf("got %v, want 2", list.length)
	}

	list.enqueue(11)

	if v := list.deque(); v != 7 {
		t.Errorf("got %v, want 7", v)
	}
	if v := list.deque(); v != 9 {
		t.Errorf("got %v, want 9", v)
	}
	if v := list.peek(); v != 11 {
		t.Errorf("got %v, want 11", v)
	}
	if v := list.deque(); v != 11 {
		t.Errorf("got %v, want 11", v)
	}
	if v := list.deque(); v != 0 {
		t.Errorf("got %v, want 0", v)
	}
	if list.length != 0 {
		t.Errorf("got %v, want 0", list.length)
	}
	
	list.enqueue(69)
	if v := list.peek(); v != 69 {
		t.Errorf("got %v, want 69", v)
	}
	if list.length != 1 {
		t.Errorf("got %v, want 1", list.length)
	}
	if v := list.deque(); v != 69 {
		t.Errorf("got %v, want 69", v)
	}

	list.enqueue(420)
	if v := list.peek(); v != 420 {
		t.Errorf("got %v, want 420", v)
	}
	if list.length != 1 {
		t.Errorf("got %v, want 1", list.length)
	}

}
