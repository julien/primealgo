package main

type lruNode[T Ordered] struct {
	value      T
	prev, next *lruNode[T]
}

type lru[K, V Ordered] struct {
	length        int
	head, tail    *lruNode[V]
	lookup        map[K]*lruNode[V]
	reverseLookup map[*lruNode[V]]K
	capacity      int
}

func newLRU[K, V Ordered](capacity int) *lru[K, V] {
	lookup := make(map[K]*lruNode[V])
	reverseLookup := make(map[*lruNode[V]]K)
	return &lru[K, V]{capacity: capacity, lookup: lookup, reverseLookup: reverseLookup}
}

func (l *lru[K, V]) update(key K, value V) {
	// Does the value exist?
	// Check the cache for existence

	// If it doesn't, we need to insert
	// - Check capacity and evict if over

	// If it does, we need to update to the front of the list
	// and update the value
	node, ok := l.lookup[key]
	if !ok {
		node := &lruNode[V]{value: value}
		l.length++
		l.preprend(node)
		l.trimCache()

		l.lookup[key] = node
		l.reverseLookup[node] = key
	} else {
		l.detach(node)
		l.preprend(node)
		node.value = value
	}
}

func (l *lru[K, V]) get(key K) V {
	var val V
	// Check the cache for existence
	node, ok := l.lookup[key]
	if !ok {
		return val
	}

	// Update the value we found and move it to the front
	l.detach(node)
	l.preprend(node)

	// Return out the value found or nil if it doesn't exist
	return node.value
}

func (l *lru[K, V]) detach(node *lruNode[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if l.head == node {
		l.head = l.head.next
	}

	if l.tail == node {
		l.tail = l.tail.prev
	}

	node.next = nil
	node.prev = nil
}

func (l *lru[K, V]) preprend(node *lruNode[V]) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *lru[K, V]) trimCache() {
	if l.length <= l.capacity {
		return
	}

	tail := l.tail
	l.detach(l.tail)
	key := l.reverseLookup[tail]
	delete(l.lookup, key)
	delete(l.reverseLookup, tail)
	l.length--
}
