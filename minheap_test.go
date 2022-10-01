package main

import "testing"

func TestMinHeap(t *testing.T) {
	t.Parallel()

	heap := NewMinHeap[int]()

	if v := heap.Length(); v != 0 {
		t.Errorf("got %v, want 0", v)
	}

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	if v := heap.Length(); v != 8 {
		t.Errorf("got %v, want 8", v)
	}
	if v := heap.Delete(); v != 1 {
		t.Errorf("got %v, want 1", v)
	}
	if v := heap.Delete(); v != 3 {
		t.Errorf("got %v, want 3", v)
	}
	if v := heap.Delete(); v != 4 {
		t.Errorf("got %v, want 4", v)
	}
	if v := heap.Delete(); v != 5 {
		t.Errorf("got %v, want 5", v)
	}
	if v := heap.Length(); v != 4 {
		t.Errorf("got %v, want 4", v)
	}
	if v := heap.Delete(); v != 7 {
		t.Errorf("got %v, want 7", v)
	}
	if v := heap.Delete(); v != 8 {
		t.Errorf("got %v, want 8", v)
	}
	if v := heap.Delete(); v != 69 {
		t.Errorf("got %v, want 69", v)
	}
	if v := heap.Delete(); v != 420 {
		t.Errorf("got %v, want 420", v)
	}
	if v := heap.Length(); v != 0 {
		t.Errorf("got %v, want 0", v)
	}
}
