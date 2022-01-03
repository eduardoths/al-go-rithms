package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"Success to get existent storage":    testGetExistent,
		"Fails to get unexistent":            testGetUnexistent,
		"Asserts reordering when existent":   testGetAndReorderExistent,
		"Asserts reordering when unexistent": testGetAndReorderUnexistent,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testGetExistent(t *testing.T) {
	cache := FakeCache()
	input := 2
	expected := 20
	actual := cache.Get(input)

	assert.Equal(t, expected, actual)
}

func testGetUnexistent(t *testing.T) {
	cache := FakeCache()
	input := 10231
	expected := -1
	actual := cache.Get(input)

	assert.Equal(t, expected, actual)
}

func testGetAndReorderExistent(t *testing.T) {
	cache := FakeCache()
	first := cache.storage
	second := first.next
	third := second.next
	input := 2

	cache.Get(input)
	assert.Equal(t, second, cache.storage)
	assert.Equal(t, first, second.next)
	assert.Equal(t, third, first.next)

}

func testGetAndReorderUnexistent(t *testing.T) {
	cache := FakeCache()
	first := cache.storage
	second := first.next
	third := second.next
	input := 10231

	cache.Get(input)
	assert.Equal(t, first, cache.storage)
	assert.Equal(t, second, first.next)
	assert.Equal(t, third, second.next)
}
