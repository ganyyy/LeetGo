package main

func search(nums []int, target int) bool {
	// 做法1: 先找到翻转的地址

	var l, r = 0, len(nums) - 1

	var mid int
	for l <= r {
		// 首先去除两边的重复数字.. 跪在这一步上了... 草
		for l < r && nums[l] == nums[l+1] {
			l++
		}
		for l < r && nums[r] == nums[r-1] {
			r--
		}
		mid = l + (r-l)>>1
		if nums[mid] == target || nums[l] == target || nums[r] == target {
			return true
		}
		// if nums[mid] > target {
		//     // target < nums[mid]
		//     if nums[l] > nums[mid] {
		//         // 拐点在左半部分, 小于 target的值只能在 [k, mid-1]中, 向左查询
		//         r = mid-1
		//     } else {
		//         // 拐点在右半部分
		//         if nums[r] < target {
		//             // 如果目标值大于右边的最大值, 此时[l, mid]是有序的, 从左边查找
		//             r = mid-1
		//         } else {
		//             // 否则, 目标值应该在[k, r]中查找
		//             l = mid+1
		//         }
		//     }
		// } else {
		//     // target > nums[mid]
		//     if nums[r] < nums[mid] {
		//         // 拐点在右半部分, 大于target的值只能在 [mid+1, k-1]中查找
		//         l = mid+1
		//     } else {
		//         // 拐点在左半部分
		//         if target > nums[l] {
		//             // 如果目标值大于 左边的最小值, 说明目标值一定在 [l, k-1]中
		//             r = mid-1
		//         } else {
		//             // 否则, 此时[mid, r]是有序的, 走通用的二分查找即可
		//             l = mid+1
		//         }
		//     }
		// }
		if nums[mid] >= nums[l] {
			// 左半部分有序
			if target < nums[mid] && target >= nums[l] {
				// 从有序部分进行查找
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 右半部分有序
			if target > nums[mid] && target <= nums[r] {
				// 从有序部分进行查找
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
