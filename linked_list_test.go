package main

import "testing"

func TestLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(5)
	list.Append(7)
	list.Append(9)

	if v := list.Get(2); v != 9 {
		t.Errorf("got %v, want 9", v)
	}
	if v := list.RemoveAt(1); v != 7 {
		t.Errorf("got %v, want 7", v)
	}
	if v := list.Length(); v != 2 {
		t.Errorf("got %v, want 2", v)
	}

	list.Append(11)
	if v := list.RemoveAt(1); v != 9 {
		t.Errorf("got %v, want 9", v)
	}
	if v := list.Remove(9); v != 0 {
		t.Errorf("got %v, want 0", v)
	}
	if v := list.RemoveAt(0); v != 5 {
		t.Errorf("got %v, want 5", v)
	}
	if v := list.RemoveAt(0); v != 11 {
		t.Errorf("got %v, want 11", v)
	}
	if v := list.Length(); v != 0 {
		t.Errorf("got %v, want 0", v)
	}

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	if v := list.Get(2); v != 5 {
		t.Errorf("got %v, want 5", v)
	}
	if v := list.Get(0); v != 9 {
		t.Errorf("got %v, want 9", v)
	}
	if v := list.Remove(9); v != 9 {
		t.Errorf("got %v, want 9", v)
	}
	if v := list.Length(); v != 2 {
		t.Errorf("got %v, want 2", v)
	}
	if v := list.Get(0); v != 7 {
		t.Errorf("got %v, want 7", v)
	}

	list.InsertAt(3, 0)
	if v := list.Get(0); v != 3 {
		t.Errorf("got %v, want 3", v)
	}
	if v := list.Length(); v != 3 {
		t.Errorf("got %v, want 3", v)
	}

	list.InsertAt(4, 1)
	if v := list.Get(1); v != 4 {
		t.Errorf("got %v, want 4", v)
	}
	if v := list.Length(); v != 4 {
		t.Errorf("got %v, want 4", v)
	}

	list.InsertAt(72, 4)
	if v := list.Get(4); v != 72 {
		t.Errorf("got %v, want 72", v)
	}
}
