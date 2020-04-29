package main

import "fmt"

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

type MountainArray struct {
	arr []int
}

func (mountain *MountainArray) get(index int) int {
	return mountain.arr[index]
}
func (mountain *MountainArray) length() int {
	return len(mountain.arr)
}

func NewMountainArray(arr []int) *MountainArray {
	return &MountainArray{arr: arr}
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	cache := make(map[int]int)
	var t int
	get := func(index int) int {
		if v, ok := cache[index]; ok {
			return v
		} else {
			t = mountainArr.get(index)
			cache[index] = t
			return t
		}
	}

	var top int
	ln := mountainArr.length()
	left, right := 0, ln-1

	// 找到顶点top的index
	for left < right {
		mid := left + (right-left)>>1
		r, l, m := get(mid+1), get(mid-1), get(mid)
		if m > r && m > l {
			top = mid
			break
		} else if l < m && m < r {
			left = mid
		} else {
			right = mid
		}
	}

	if target == get(top) {
		return top
	}
	// 找左边
	left, right = 0, top-1
	for left <= right {
		mid := left + (right-left)>>1
		v := get(mid)
		if v == target {
			return mid
		}
		if v > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 找右边
	left, right = top+1, ln-1
	for left <= right {
		mid := left + (right-left)>>1
		v := get(mid)
		if v == target {
			return mid
		}
		if v < target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	/**
	[1,5,2]
	1
	*/
	mountain := NewMountainArray([]int{1, 5, 2})

	fmt.Println(findInMountainArray(5, mountain))
}
