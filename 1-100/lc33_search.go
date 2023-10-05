package main

import (
	"fmt"
)

func search33_1(nums []int, target int) int {
	ln := len(nums)
	if ln == 0 {
		return -1
	}
	if ln == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// 二分找
	left, right, index := 0, ln-1, -1
	for left < right {
		if nums[left] == target {
			index = left
			break
		}
		if nums[right] == target {
			index = right
			break
		}
		// 看中值
		mid := (left + right) / 2
		if nums[mid] == target {
			index = mid
			break
		}

		if nums[mid] < nums[right] {
			// 右边升序
			if nums[mid] < target && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 左边升序
			if nums[mid] > target && target > nums[left] {
				right = mid - 1
			} else {
				left = mid
			}
		}
	}
	return index
}

func search33_2(nums []int, target int) int {
	// 怎么找到旋转点呢?

	// 肯定还是二分
	// 假设中点在 旋转点的左侧
	// 4,5,6,7,0,1,2
	// 中点为7, 大于左侧和右侧. 如果查找的数字小于中点, 且大于左侧, 则向左靠;
	//                        如果查找的数字大于中点, 则向右靠
	// 假设中点在 旋转点的右侧
	// 6,7,0,1,2,4,5
	// 中点为1, 小于左侧和右侧. 如果查找的数字小于中点, 那么就往左靠
	//                        如果查找的数字大于中点, 且小于右侧, 则向右靠

	var l, r = 0, len(nums)

	for l < r {
		mid := l + (r-l)/2
		mv := nums[mid]
		if mv == target {
			return mid
		}

		lv := nums[l]
		rv := nums[r-1]

		// fmt.Printf("[%d]%d, [%d]%d, [%d][%d]\n", l,lv,r,rv,mid,mv)

		if mv > lv && mv > rv {
			// 拐点在右侧
			if target < mv && target >= lv {
				r = mid
			} else {
				l = mid + 1
			}
		} else {
			// 拐点在右侧
			if target > mv && target <= rv {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	return -1
}

func search33_3(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if r < 0 || (r == 0 && nums[0] != target) {
		return -1
	}
	// 没图说个锤子
	/*
		拐点在右侧的情况: 中点肯定是大于左侧和右侧的

						7
					⑥
				5
			4
		3
								2
							1


		拐点在左侧的情况: 中点肯定是小于左侧和右侧的

			7
		6
								5
							4
						3
					②
				1
	*/

	for l <= r {
		mid := l + (r-l)/2
		mv := nums[mid]
		if mv == target {
			return mid
		}
		// 判断是不是有拐点
		lv, rv := nums[l], nums[r]

		if mv > lv && mv > rv {
			// 拐点在中点的右侧, 此时[l,m]是有序的
			if target < mv && target >= lv {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 拐点在中点的左侧, 此时[m,r]是有序的
			if target > mv && target <= rv {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	var nums = []int{2, 4, 5, 6, 7, 0, 1}
	fmt.Println(search33_1(nums, 7))
}
