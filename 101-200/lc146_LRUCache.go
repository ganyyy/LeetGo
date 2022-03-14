package main

import (
	"container/list"
	"fmt"
	"strings"
)

type node struct {
	val  int
	key  int
	prev *node
	next *node
}

type LRUCacheOld struct {
	m    map[int]*node
	head *node
	tail *node
	c    int
}

func ConstructorOld(capacity int) LRUCacheOld {
	return LRUCacheOld{
		m:    make(map[int]*node, capacity),
		head: nil,
		tail: nil,
		c:    capacity,
	}
}

func (cache *LRUCacheOld) get(key int) int {
	if v, ok := cache.m[key]; ok {
		cache.moveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (cache *LRUCacheOld) moveToHead(v *node) {
	if v.prev != nil {
		if v == cache.tail {
			cache.tail = v.prev
		}
		v.prev.next = v.next
		if nil != v.next {
			v.next.prev = v.prev
		}
		v.next = cache.head
		v.prev = nil
		cache.head.prev = v
		cache.head = v
	}
}

func (cache *LRUCacheOld) put(key int, value int) {
	if v, ok := cache.m[key]; ok {
		v.val = value
		cache.moveToHead(v)
	} else {
		n := &node{
			val:  value,
			key:  key,
			prev: nil,
			next: cache.head,
		}
		cache.m[key] = n
		if nil != cache.head {
			cache.head.prev = n
		}
		cache.head = n
		if cache.tail == nil {
			cache.tail = n
		}
		// 找到尾节点, 删除
		if len(cache.m) > cache.c {
			t := cache.tail
			cache.tail = t.prev
			cache.tail.next = nil
			delete(cache.m, t.key)
		}
	}
}

type Node struct {
	Key, Val int
}

type LRUCache struct {
	keys     map[int]*list.Element
	list     *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keys:     make(map[int]*list.Element, capacity),
		list:     list.New(),
		capacity: capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	if ele, ok := l.keys[key]; ok {
		l.list.MoveToFront(ele)
		return ele.Value.(Node).Val
	} else {
		return -1
	}
}

func (l *LRUCache) Put(key int, value int) {
	if ele, ok := l.keys[key]; ok {
		ele.Value = Node{
			Key: key, Val: value,
		}
		l.list.MoveToFront(ele)
	} else {
		if len(l.keys) >= l.capacity {
			var val = l.list.Remove(l.list.Back())
			delete(l.keys, val.(Node).Key)
		}
		l.keys[key] = l.list.PushFront(Node{
			Key: key,
			Val: value,
		})
	}
}

func (l *LRUCache) Show() {
	var sb strings.Builder
	for head := l.list.Front(); head != nil; head = head.Next() {
		sb.WriteString(fmt.Sprintf("%+v,", head.Value))
	}
	fmt.Println(sb.String())
}
