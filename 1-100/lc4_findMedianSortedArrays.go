//go:build ignore

package main

import (
	"container/heap"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 保证 nums1是短的那一个
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	// 如果有一个为空就可以提前返回
	len1, len2 := len(nums1), len(nums2)
	val1, val2 := 0, 0
	if len1 == 0 {
		mid := len2 / 2
		if len2&1 == 0 {
			return float64(nums2[mid]+nums2[mid-1]) / 2
		} else {
			return float64(nums2[mid])
		}
	}
	// 确定两个数组的起点
	start1, start2 := 0, 0
	// 确定两个数组合并后的中点位置
	mid := (len1 + len2 + 1) / 2
	// 获取数组指定位置的值, 如果越界就返回最后一个
	getVal := func(arr []int, pos int) int {
		if pos >= len(arr) {
			pos = len(arr) - 1
		}
		return arr[pos]
	}
	// 获取两个位置中的较小的值
	getMin := func(pos1, pos2 int) int {
		if pos1 < 0 || pos1 >= len1 {
			return nums2[pos2]
		}
		if pos2 < 0 || pos2 >= len2 {
			return nums1[pos1]
		}
		if nums1[pos1] < nums2[pos2] {
			return nums1[pos1]
		}
		return nums2[pos2]
	}

	for {
		// 跳出条件1: nums1数组到头了
		// 这种情况是 nums1的最大值小于nums2的中间值
		// 很明显, 中值应该在nums2中
		if start1 >= len1 {
			// 此时中值1取nums2的值
			val1 = nums2[start2+mid-1]
			// 中值2取两个位置中的小值
			val2 = getMin(start1, start2+mid)
			break
		}
		// 跳出条件2: mid==1
		// 此时中值分布在两个数组中
		if mid == 1 {
			// 以小的值为基准, 中值1一定是二者最小的值
			// 中值2从中值1位置的下一个和另一数组的start取最小值
			if nums1[start1] <= nums2[start2] {
				val1 = nums1[start1]
				val2 = getMin(start1+1, start2)
			} else {
				val1 = nums2[start2]
				val2 = getMin(start1, start2+1)
			}
			break
		}
		// 每次排除 mid / 2个元素
		k := mid / 2
		// 剩余的mid
		mid -= k
		// 比较两者对应第k位数字的大小
		if getVal(nums1, start1+k-1) > getVal(nums2, start2+k-1) {
			start2 += k
		} else {
			// 因为nums1的数组长度是最小的, 所以要判断一下start1是否越界
			// 如果越界了说明nums1的最大值小于 nums2的指定位置的值
			// 为了保证mid的正确取值, 需要将多去的再加回来
			start1 += k
			if start1 > len1 {
				mid += start1 - len1
			}
		}
	}

	if (len1+len2)&1 == 0 {
		return float64(val1+val2) / 2
	} else {
		return float64(val1)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	// left: nums1的头
	// right: nums1的尾
	// mid: nums1和nums2的中位数中偏右的那个位置

	heap.Fix()

	/*
		相当于找到nums1中的一个位置i, 和nums2中的一个位置j
		首先: i+j = (len1+len2+1)/2. 这个相当于找的是中位数偏右的那个位置, 奇数情况下-1就是中位数, 偶数情况下-1就是左边的那个数
		然后需要满足的是: nums1[i-1] <= nums2[j] && nums2[j-1] <= nums1[i]
		在此条件下针对较长的数组进行二分查找, 逼近i的位置
		如果 nums1[i-1] > nums2[j], 说明i太大了, 需要减小; 如果 nums2[j-1] > nums1[i], 说明i太小了, 需要增大
	*/

	len1, len2 := len(nums1), len(nums2)
	left, right, mid := 0, len1, (len1+len2+1)/2
	for left <= right {
		// 原理是通过依次向中间逼近找到

		// 每次都取nums1的中位
		i := (left + right) / 2
		// 相应的, mid-i就是nums2的中位?
		j := mid - i
		if i < right && nums1[i] < nums2[j-1] {
			// i太小了, 需要增大
			left = i + 1
		} else if i > left && nums1[i-1] > nums2[j] {
			// i太大了, 需要减小
			right = i - 1
		} else {
			// 到这里, 就可以保证: nums1[i-1] <= nums2[j] && nums2[j-1] <= nums1[i]
			// 相当于要从 nums1[i-1], nums1[i], nums2[j-1], nums2[j]中找到中位数

			// left是左边的最大值
			var left int
			if i == 0 {
				left = nums2[j-1]
			} else if j == 0 {
				left = nums1[i-1]
			} else {
				left = max(nums1[i-1], nums2[j-1])
			}

			// 长度为奇数, 返回left
			if (len1+len2)&1 == 1 {
				return float64(left)
			}

			// right是右边的最小值
			var right int
			if i == len1 {
				right = nums2[j]
			} else if j == len2 {
				right = nums1[i]
			} else {
				right = min(nums1[i], nums2[j])
			}

			// 长度为偶数, 返回(left+right)/2
			return float64(left+right) / 2
		}
	}
	return 0
}

func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	// 取两个数组中较短的为 nums1
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	// 保留两者的长度
	m, n := len(nums1), len(nums2)

	// 计算总偏移量 防溢出写法
	var total = m + (n-m+1)>>1

	left, right := 0, m

	// i, j 分别表示 两个数组分割线的右边的第一个数的下标
	// 1, 2, 3 |, 4, 5
	// 1, 2, | 3, 4, 5
	// 此时 i = 3, j = 2. 满足 i+j = total
	// 同时还需要满足 nums1[i-1] <= nums2[j] && nums2[j-1] <= nums1[i]
	// 即分割线左边的都要小于分割线右边的(本身nums1[i-1]<nums1[i], nums2[j-1]<nums2[j])
	for left < right {
		i := left + (right-left+1)>>1
		j := total - i
		// 凡是取i-1之类的,要注意边界问题
		if nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			left = i
		}
	}
	i, j := left, total-left

	nums1LeftMax := math.MinInt32
	if i != 0 {
		nums1LeftMax = nums1[i-1]
	}
	nums1RightMin := math.MaxInt32
	if i != m {
		nums1RightMin = nums1[i]
	}
	nums2LeftMax := math.MinInt32
	if j != 0 {
		nums2LeftMax = nums2[j-1]
	}
	nums2RightMin := math.MaxInt32
	if j != n {
		nums2RightMin = nums2[j]
	}

	// 区分偶数和奇数个 个数
	if (m+n)&1 == 1 {
		return float64(getMax(nums1LeftMax, nums2LeftMax))
	} else {
		return float64(getMax(nums1LeftMax, nums2LeftMax)+getMin(nums1RightMin, nums2RightMin)) / 2
	}
}

func getMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func getMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {

}
