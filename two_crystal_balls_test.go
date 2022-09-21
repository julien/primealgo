package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestTwoCrystalBalls(t *testing.T) {
	t.Parallel()

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	idx := int(math.Floor(float64(r.Intn(1000))))
	data := make([]bool, 1000)
	for i := 0; i < len(data); i++ {
		data[i] = false
	}
	for i := idx; i < len(data); i++ {
		data[i] = true
	}

	if result := twoCrystalBalls(data); result != idx {
		t.Errorf("got %v, want %v", result, idx)
	}

	data = make([]bool, 1000)
	for i := 0; i < len(data); i++ {
		data[i] = false
	}

	if result := twoCrystalBalls(data); result != -1 {
		t.Errorf("got %v, want %v", result, -1)
	}
}
