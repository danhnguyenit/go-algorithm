package main

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}
type LRUCache struct {
	store    map[int]*node
	capacity int
	head     *node
	tail     *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, store: make(map[int]*node)}
}
func (this *LRUCache) removeNode(x *node) {
	if x == this.head {
		this.head = x.next
	}
	if x == this.tail {
		this.tail = x.prev
	}
	if x.next != nil {
		x.next.prev = x.prev
	}
	if x.prev != nil {
		x.prev.next = x.next
	}
}
func (this *LRUCache) appendFirst(x *node) {
	x.prev = nil
	x.next = this.head
	if x.next != nil {
		x.next.prev = x
	}
	this.head = x
	if this.tail == nil {
		this.tail = x
	}
}

func (this *LRUCache) Get(key int) int {
	x, ok := this.store[key]
	if !ok {
		return -1
	}
	this.removeNode(x)
	this.appendFirst(x)
	return x.value
}

func (this *LRUCache) Put(key int, value int) {
	x, ok := this.store[key]
	if !ok {
		x = &node{key: key, value: value}
		this.store[key] = x
	} else {
		x.value = value
		this.removeNode(x)
	}
	this.appendFirst(x)
	if len(this.store) > this.capacity {
		x = this.tail
		if x != nil {
			this.removeNode(x)
			delete(this.store, x.key)
		}
	}
}
