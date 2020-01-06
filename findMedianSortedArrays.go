package main

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

func m2(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 < len2 {
		nums1, nums2 = nums2, nums1
		len1, len2 = len2, len1
	}
	// head: nums1的头
	// tail: nums1的尾
	// mid: nums1和nums2的中位
	head, tail, mid := 0, len1, (len1+len2+1)/2
	for head < tail {

		// 原理是通过依次向中间逼近找到
		// 感觉这个时间复杂度不对啊
		// 应该不是O(log n)

		// 每次都取nums1的中位
		i := (head + tail) / 2
		// 相应的, mid-i就是nums2的中位?
		j := mid - i
		if i < tail && nums1[i] < nums2[j-1] {
			//
			head = i + 1
		} else if i > head && nums1[i-1] > nums2[j] {
			tail = i - 1
		} else {
			var left int
			if i == 0 {
				left = nums2[j-1]
			} else if j == 0 {
				left = nums1[i-1]
			} else {
				left = max(nums1[i-1], nums2[j-1])
			}
			if (len1+len2)&1 == 1 {
				return float64(left)
			}

			var right int
			if i == len1 {
				right = nums2[j]
			} else if j == len2 {
				right = nums1[i]
			} else {
				right = min(nums1[i], nums2[j])
			}
			return float64(left+right) / 2
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
