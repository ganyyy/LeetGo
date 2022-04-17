//go:build ignore

package main

import "math/rand"

var position = make(map[int]int)
var arr []int

type RandomizedSet struct {
	position map[int]int
	arr      []int
}

func Constructor() RandomizedSet {
	for k := range position {
		delete(position, k)
	}
	return RandomizedSet{
		position: position,
		arr:      arr[:0],
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.position[val]; ok {
		return false
	}
	this.position[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.position[val]; !ok {
		return false
	} else {
		var lastIdx = len(this.arr) - 1
		var last = this.arr[lastIdx]
		this.arr = this.arr[:lastIdx]
		if last != val {
			this.arr[idx] = last
			this.position[last] = idx
		}
		delete(this.position, val)
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.arr) == 0 {
		return 0
	}
	return this.arr[rand.Intn(len(this.arr))]
}
