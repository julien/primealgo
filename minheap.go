package main

type MinHeap[T Ordered] struct {
	length int
	data   []T
}

func NewMinHeap[T Ordered]() *MinHeap[T] {
	return &MinHeap[T]{data: make([]T, 0), length: 0}
}

func (m *MinHeap[T]) Insert(value T) {
	m.data = append(m.data, value)
	m.heapifyUp(m.length)
	m.length++
}

func (m *MinHeap[T]) Delete() T {
	var out T
	if m.length == 0 {
		return out
	}

	out = m.data[0]
	m.length--

	if m.length == 0 {
		m.data = make([]T, 0)
		return out
	}

	m.data[0] = m.data[m.length]
	m.heapifyDown(0)
	return out
}

func (m *MinHeap[T]) heapifyDown(idx int) {
	lIdx := m.leftChild(idx)
	rIdx := m.rightChild(idx)

	if idx >= m.length || lIdx >= m.length {
		return
	}

	lv := m.data[lIdx]
	rv := m.data[rIdx]
	v := m.data[idx]

	if lv > rv && v > rv {
		m.data[idx] = rv
		m.data[rIdx] = v
		m.heapifyDown(rIdx)
	} else if rv > lv && v > lv {
		m.data[idx] = lv
		m.data[lIdx] = v
		m.heapifyDown(lIdx)
	}
}

func (m *MinHeap[T]) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := m.parent(idx)
	pv := m.data[p]
	v := m.data[idx]

	if pv > v {
		m.data[idx] = pv
		m.data[p] = v
		m.heapifyUp(p)
	}
}

func (m MinHeap[T]) Length() int {
	return m.length
}

func (m MinHeap[T]) parent(idx int) int {
	return (idx - 1) / 2
}

func (m MinHeap[T]) leftChild(idx int) int {
	return idx*2 + 1
}

func (m MinHeap[T]) rightChild(idx int) int {
	return idx*2 + 2
}
