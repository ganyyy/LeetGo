package main

import "fmt"

// 归并排序的方式解决问题
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}

	// 设前缀和数组 preSum, 等同于求得 preSum[i]-preSum[j] ∈ [lower, upper] 的个数

	var mergeCnt func([]int) int

	mergeCnt = func(arr []int) int {
		var n = len(arr)
		if n <= 1 {
			return 0
		}
		// 将nums分为左右两半
		var n1 = make([]int, n/2)
		copy(n1, arr[:n/2])
		var n2 = make([]int, n-n/2)
		copy(n2, arr[n/2:])

		// 先对左右两个区间进行归并, 使其满足有序的条件, 同时获得满足要求的个数
		// 区间的排序并不影响相对位置, n2所有的数字都在n1之后. 所以可以继续计算满足条件的个数
		n = mergeCnt(n1) + mergeCnt(n2)

		// l, r 表示 n2[l:r] + n1[i] 属于 [lower, upper]
		// 此时n1, n2都是有序的
		var l, r int
		// 对于每一个n1中的数字v, 满足 v + each n2[l:r] ∈ [lower, upper], 即是合法的区间和
		// 直接加到总的计数上
		for _, v := range n1 {
			for l < len(n2) && n2[l]-v < lower {
				l++
			}
			for r < len(n2) && n2[r]-v <= upper {
				r++
			}
			// 加上当前区间的值
			n += r - l
		}
		// 执行归并. l, r表示两个数组的当前指针
		l, r = 0, 0
		for i := range arr {
			if l < len(n1) && (r == len(n2) || n1[l] <= n2[r]) {
				arr[i] = n1[l]
				l++
			} else {
				arr[i] = n2[r]
				r++
			}
		}
		return n
	}

	var preFixSum = make([]int, len(nums)+1)
	for i, v := range nums {
		preFixSum[i+1] = preFixSum[i] + v
	}
	return mergeCnt(preFixSum)
}

// 暴力解法不可取
func countRangeSum2(nums []int, lower int, upper int) int {
	if len(nums) == 0 {
		return 0
	}
	var res int
	var sum int
	// 外层循环表示 选取的起点
	for i := 0; i < len(nums); i++ {
		// 内层循环表示累加和
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= lower && sum <= upper {
				res++
			}
		}
		sum = 0
	}
	return res
}

func mergeSort(nums []int) {
	if len(nums) == 0 {
		return
	}
	var tmp = make([]int, len(nums))
	var helper func([]int)

	helper = func(nums []int) {
		var ln = len(nums)
		if ln <= 1 {
			return
		}
		var mid = ln / 2
		helper(nums[:mid])
		helper(nums[mid:])
		var p1, p2 int
		for i := 0; i < len(nums); i++ {
			if p1 < mid && (p2 == ln-mid || nums[p1] < nums[mid+p2]) {
				tmp[i] = nums[p1]
				p1++
			} else {
				tmp[i] = nums[mid+p2]
				p2++
			}
		}
		for i := 0; i < len(nums); i++ {
			nums[i] = tmp[i]
		}
	}

	helper(nums)
}

func main() {
	var t = []int{1, 1231, 41, 2341, 3, 14, 1, 523, 5, 23, 231, 3, 1245, 1, 412, 41, 41}
	mergeSort(t)
	fmt.Println(t)
}
