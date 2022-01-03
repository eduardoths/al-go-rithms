package lrucache

func FakeCache() *LRUCache {
	first := &Storage{
		key:   1,
		value: 10,
	}
	second := &Storage{
		key:   2,
		value: 20,
	}
	third := &Storage{
		key:   3,
		value: 30,
	}
	first.next = second
	second.next = third

	return &LRUCache{
		capacity: 3,
		storage:  first,
	}
}
