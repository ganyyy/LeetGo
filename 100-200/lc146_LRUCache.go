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

type ListNode struct {
	Key, Val int
}

type LRUCache2 struct {
	keys     map[int]*list.Element
	list     *list.List
	capacity int
}

func Constructor146(capacity int) LRUCache2 {
	return LRUCache2{
		keys:     make(map[int]*list.Element, capacity),
		list:     list.New(),
		capacity: capacity,
	}
}

func (l *LRUCache2) Get(key int) int {
	if ele, ok := l.keys[key]; ok {
		l.list.MoveToFront(ele)
		return ele.Value.(*ListNode).Val
	} else {
		return -1
	}
}

func (l *LRUCache2) Put(key int, value int) {
	if ele, ok := l.keys[key]; ok {
		ele.Value = &ListNode{
			Key: key, Val: value,
		}
		l.list.MoveToFront(ele)
	} else {
		if len(l.keys) >= l.capacity {
			var val = l.list.Remove(l.list.Back())
			delete(l.keys, val.(*ListNode).Key)
		}
		l.keys[key] = l.list.PushFront(&ListNode{
			Key: key,
			Val: value,
		})
	}
}

func (l *LRUCache2) Show() {
	var sb strings.Builder
	for head := l.list.Front(); head != nil; head = head.Next() {
		sb.WriteString(fmt.Sprintf("%+v,", head.Value))
	}
	fmt.Println(sb.String())
}

type Node146_2 struct {
	Key, Val int
}

func To(node *list.Element) *Node146_2 {
	return node.Value.(*Node146_2)
}

type LRUCache struct {
	allNode  *list.List
	nodes    map[int]*list.Element
	capacity int
}

func Constructor146_2(capacity int) LRUCache {
	return LRUCache{
		allNode:  list.New(),
		nodes:    make(map[int]*list.Element, capacity),
		capacity: capacity,
	}
}

func (cache *LRUCache) Get(key int) int {
	if element, ok := cache.nodes[key]; ok {
		cache.allNode.MoveToFront(element)
		return To(element).Val
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	// 存在
	if element, ok := cache.nodes[key]; ok {
		To(element).Val = value
		cache.allNode.MoveToFront(element)
		return
	}
	if cache.capacity <= cache.allNode.Len() {
		// 空间已满
		last := cache.allNode.Back()
		cache.allNode.Remove(last)
		delete(cache.nodes, To(last).Key)
	}
	// 插入新节点
	node := &Node146_2{
		Key: key, Val: value,
	}
	cache.nodes[key] = cache.allNode.PushFront(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor146_2(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
