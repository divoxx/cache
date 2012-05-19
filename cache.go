package cache

type Key interface{}

type Value interface{}

type node struct {
	Key
	Value
}

type Interface interface {
	Put(Key, Value)
	Get(Key) Value
}
