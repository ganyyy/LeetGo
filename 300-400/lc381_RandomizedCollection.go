package main

import (
	"math/rand"
)

type RandomizedCollection struct {
	idx  map[int]map[int]struct{}
	nums []int
}

var empty = struct{}{}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		idx: make(map[int]map[int]struct{}),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	var ids, ok = this.idx[val]
	if !ok {
		ids = make(map[int]struct{})
		this.idx[val] = ids
	}
	ids[len(this.nums)] = empty
	this.nums = append(this.nums, val)
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	var ids, ok = this.idx[val]
	if !ok {
		return false
	}
	// 获取当前值的一个索引
	var i int
	for i = range ids {
		break
	}
	// 去掉索引
	delete(ids, i)

	// 维护 nums 数组, 将当前要删除的位置 替换到 数组末尾
	var n = len(this.nums) - 1
	this.nums[i] = this.nums[n]

	// 去掉替换位置的旧索引
	delete(this.idx[this.nums[i]], n)
	if i < n {
		// 将当前位置的索引加入进去
		this.idx[this.nums[i]][i] = empty
	}
	// 检查当前元素 是否还有剩余元素
	if len(ids) == 0 {
		delete(this.idx, val)
	}

	// 删除的nums 元素
	this.nums = this.nums[:n]
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	var r = Constructor()
	r.Insert(1)
	r.Remove(2)
	r.Insert(2)
	r.GetRandom()
	r.Remove(1)
	r.Insert(2)
	r.GetRandom()
}
