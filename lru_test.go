package main

import "testing"

func TestLRU(t *testing.T) {
	t.Parallel()

	l := newLRU[string, int](3)

	if v := l.get("foo"); v != 0 {
		t.Errorf("got %v, want 0", v)
	}
	l.update("foo", 69)
	if v := l.get("foo"); v != 69 {
		t.Errorf("got %v, want 69", v)
	}

	l.update("bar", 420)
	if v := l.get("bar"); v != 420 {
		t.Errorf("got %v, want 420", v)
	}

	l.update("baz", 1337)
	if v := l.get("baz"); v != 1337 {
		t.Errorf("got %v, want 1337", v)
	}

	l.update("ball", 69420)
	if v := l.get("ball"); v != 69420 {
		t.Errorf("got %v, want 69420", v)
	}
	if v := l.get("foo"); v != 0 {
		t.Errorf("got %v, want 0", v)
	}
	if v := l.get("bar"); v != 420 {
		t.Errorf("got %v, want 420", v)
	}
	l.update("foo", 69)
	if v := l.get("bar"); v != 420 {
		t.Errorf("got %v, want 420", v)
	}
	if v := l.get("foo"); v != 69 {
		t.Errorf("got %v, want 69", v)
	}
	if v := l.get("baz"); v != 0 {
		t.Errorf("got %v, want 0", v)
	}

	l.update("foo", 420)
	if v := l.get("foo"); v != 420 {
		t.Errorf("got %v, want 420", v)
	}
}
