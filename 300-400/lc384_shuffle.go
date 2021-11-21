package main

import "math/rand"

type Solution struct {
	Nums  []int
	SNums []int
}

func Constructor(nums []int) Solution {
	var tmp = make([]int, len(nums))
	copy(tmp, nums)
	return Solution{
		Nums:  nums,
		SNums: tmp,
	}
}

func (this *Solution) Reset() []int {
	copy(this.SNums, this.Nums)
	return this.Nums
}

func (this *Solution) Shuffle() []int {
	rand.Shuffle(len(this.SNums), func(i, j int) {
		this.SNums[i], this.SNums[j] = this.SNums[j], this.SNums[i]
	})
	return this.SNums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
