package cache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	c := NewLRU(3)

	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)

	if v := c.Get("a"); v != 1 {
		t.Error("Cached value of a should have been 2 but it was", v)
	}

	if v := c.Get("c"); v != 3 {
		t.Error("Cached value of c should have been 4 but it was", v)
	}

	c.Put("d", 4)

	if v := c.Get("b"); v != nil {
		t.Error("Cache lookup for b should be a miss")
	}
}
