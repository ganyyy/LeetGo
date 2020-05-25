package main

type node struct {
	val  int
	key  int
	prev *node
	next *node
}

type LRUCache struct {
	m    map[int]*node
	head *node
	tail *node
	c    int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:    make(map[int]*node, capacity),
		head: nil,
		tail: nil,
		c:    capacity,
	}
}

func (this *LRUCache) get(key int) int {
	if v, ok := this.m[key]; ok {
		this.moveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (this *LRUCache) moveToHead(v *node) {
	if v.prev != nil {
		if v == this.tail {
			this.tail = v.prev
		}
		v.prev.next = v.next
		if nil != v.next {
			v.next.prev = v.prev
		}
		v.next = this.head
		v.prev = nil
		this.head.prev = v
		this.head = v
	}
}

func (this *LRUCache) put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.val = value
		this.moveToHead(v)
	} else {
		n := &node{
			val:  value,
			key:  key,
			prev: nil,
			next: this.head,
		}
		this.m[key] = n
		if nil != this.head {
			this.head.prev = n
		}
		this.head = n
		if this.tail == nil {
			this.tail = n
		}
		// 找到尾节点, 删除
		if len(this.m) > this.c {
			t := this.tail
			this.tail = t.prev
			this.tail.next = nil
			delete(this.m, t.key)
		}
	}
}

func main() {
	cache := Constructor(3 /* 缓存容量 */)

	cache.put(1, 1)
	cache.put(2, 2)
	cache.get(1)    // 返回  1
	cache.put(3, 3) // 该操作会使得关键字 2 作废
	cache.get(2)    // 返回 -1 (未找到)
	cache.put(4, 4) // 该操作会使得关键字 1 作废
	cache.get(1)    // 返回 -1 (未找到)
	cache.get(3)    // 返回  3
	cache.get(4)    // 返回  4
}
