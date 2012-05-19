package cache

import (
	"container/list"
)

type LRU struct {
	size int
	list *list.List
	indx map[Key]*list.Element
}

func NewLRU(size int) *LRU {
	return &LRU{
		size: size,
		list: list.New(),
		indx: make(map[Key]*list.Element),
	}
}

func (c *LRU) Put(k Key, v Value) {
	newNode := &node{k, v}

	if curr := c.indx[k]; curr != nil {
		c.list.Remove(curr)

	} else {
		if c.list.Len() == c.size {
			c.expire()
		}
	}

	c.indx[k] = c.list.PushFront(newNode)
}

func (c *LRU) Get(k Key) Value {
	if curr := c.indx[k]; curr != nil {
		c.list.MoveToFront(curr)
		return nodeFor(curr).Value
	}

	return nil
}

func (c *LRU) expire() {
	older := c.list.Back()
	olderKey := nodeFor(older).Key
	delete(c.indx, olderKey)
	c.list.Remove(older)
}

func nodeFor(el *list.Element) *node {
	return el.Value.(*node)
}
