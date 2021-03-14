package main

import "container/list"

type MyHashMapOld struct {
	vals []int32
}

/** Initialize your data structure here. */
func Constructor() MyHashMapOld {
	return MyHashMapOld{
		// 这是一个脑淤血的写法. 但是速度为啥这么差呢...
		vals: make([]int32, 1e6+1),
	}
}

/** value will always be non-negative. */
func (this *MyHashMapOld) Put(key int, value int) {
	this.vals[key] = int32(value + 1)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMapOld) Get(key int) int {
	if this.vals[key] == 0 {
		return -1
	}
	return int(this.vals[key] - 1)
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMapOld) Remove(key int) {
	this.vals[key] = 0
}

/**
 * Your MyHashMapOld object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

const base = 769

type entry struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}

func Constructor2() MyHashMap {
	return MyHashMap{make([]list.List, base)}
}

func (m *MyHashMap) hash(key int) int {
	return key % base
}

func (m *MyHashMap) Put(key, value int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	m.data[h].PushBack(entry{key, value})
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			m.data[h].Remove(e)
		}
	}
}
