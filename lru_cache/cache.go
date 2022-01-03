package lrucache

type LRUCache struct {
	capacity int
	storage  *Storage
}

type Storage struct {
	key   int
	value int
	next  *Storage
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
	}
}

func (lc *LRUCache) clearOverflow() {
	current := lc.storage
	counter := 1
	for current != nil {
		if counter == lc.capacity || current == nil {
			current.next = nil
			return
		}
		counter++
		current = current.next
	}
}

func (lc *LRUCache) delete(key int) {
	first := lc.storage
	current := lc.storage
	last := current
	for current != nil {
		if current.key == key {
			if first == current {
				lc.storage = lc.storage.next
			}
			last.next = current.next
			return
		}
		last = current
		current = current.next
	}
}

func (lc *LRUCache) Get(key int) int {
	first := lc.storage
	current := lc.storage
	last := current
	for current != nil {
		if current.key == key {
			last.next = current.next
			lc.storage = current
			if first != current {
				current.next = first
			}
			return current.value
		}
		last = current
		current = current.next
	}
	return -1
}

func (lc *LRUCache) Put(key int, value int) {
	lc.delete(key)
	next := &Storage{
		key:   key,
		value: value,
		next:  lc.storage,
	}
	lc.storage = next
	lc.clearOverflow()
}
