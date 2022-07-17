//go:build ignore

package main

type MovingAverage struct {
	sum  float64
	size int
	arr  []float64
}

/** Initialize your data structure here. */

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.arr = append(this.arr, float64(val))
	this.sum += float64(val)
	if len(this.arr) > this.size {
		var first = this.arr[0]
		this.sum -= first
		this.arr = this.arr[1:]
	}
	return this.sum / float64(len(this.arr))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
