package main

import "fmt"

func medianSlidingWindow(nums []int, k int) []float64 {
	// 我日, 麻了
	// 暴力解法竟然过了....
	// 上来就这么搞, 真的好吗?

	// 直接计算肯定是没问题的,
	var res = make([]float64, len(nums)-k+1)

	var isSingle = k&1 != 0
	// queue 是一个单调的递增队列, 整体采用插入排序
	var queue = make([]int, 0, k)

	// 二分搜索, 如果存在返回这个值, 否则返回要插入的位置
	var binarySearch = func(idx int) int {
		var left, right = 0, len(queue)
		for left < right {
			var mid = left + (right-left)>>1
			if nums[queue[mid]] == nums[idx] {
				// 这里性能最差会调到o(n)
				if queue[mid] == idx {
					return mid
				}
				for i := mid - 1; i >= 0 && nums[queue[i]] == nums[idx]; i-- {
					if queue[i] == idx {
						return i
					}
				}
				for i := mid + 1; i < len(queue) && nums[queue[i]] == nums[idx]; i++ {
					if queue[i] == idx {
						return i
					}
				}
				// 如果找不到就返回mid就好了
				return mid
			} else if nums[queue[mid]] > nums[idx] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return left
	}

	var mid = k >> 1
	// 同时再维护一个位置到索引的数组
	var l, r int
	for ; r < len(nums); r++ {
		if len(queue) < cap(queue) {
			var idx = binarySearch(r)
			queue = append(queue, r)
			if idx != len(queue) {
				copy(queue[idx+1:], queue[idx:])
				queue[idx] = r
			}
		} else {
			// 核心是如何保证栈有序, 并且
			res[l] = float64(nums[queue[mid]])
			if !isSingle {
				res[l] = (res[l] + float64(nums[queue[mid-1]])) / 2
			}

			// l 出队, r入队, 找到合适的位置
			var idx = binarySearch(l)
			if idx != len(queue) {
				copy(queue[idx:], queue[idx+1:])
			}
			var tmp = queue
			queue = queue[:len(queue)-1]
			idx = binarySearch(r)
			queue = tmp
			if idx != len(queue) {
				copy(queue[idx+1:], queue[idx:])
				queue[idx] = r
			}
			l++
		}
	}
	// 最后收一下尾
	res[l] = float64(nums[queue[mid]])
	if !isSingle {
		res[l] = (res[l] + float64(nums[queue[mid-1]])) / 2
	}
	return res
}

func main() {
	/*
		[6,5,9,5,4,9,1,7,5,5]
		4
	*/
	fmt.Println(medianSlidingWindow([]int{6, 5, 9, 5, 4, 9, 1, 7, 5, 5}, 4))
}
