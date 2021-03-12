package main

type MyHashSet struct {
	// bit set
	set []uint8
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		// +1 是为了处理 边界值
		set: make([]uint8, 100000/8+1),
	}
}

func (s *MyHashSet) Add(key int) {
	s.set[key>>3] |= 1 << (key % 8)
}

func (s *MyHashSet) Remove(key int) {
	s.set[key>>3] &^= 1 << (key % 8)
}

/** Returns true if this set contains the specified element */
func (s *MyHashSet) Contains(key int) bool {
	return s.set[key>>3]&(1<<(key%8)) != 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
