package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][3]int {
	// 先排序
	sort.Ints(nums)
	numLen := len(nums)
	res := make([][3]int, 0, numLen/3)
	// 从头开始, 当前索引为k, 选取 k后的i(k+1),j(len-1)为双指针, 向中间移动
	// 1.当 nums[k] > 0, break.因为数组已排完序, 往后不可能 出现相加=0的情况
	// 2.当 nums[k] == nums[k-1](k-1 >= 0)时跳过, 因为会重复
	// 3.nums[k] + nums[i] + nums[j] = s
	//		I.  当 s > 0; j--, 并跳过重复的nums[j]
	//		II. 当 s < 0; i++, 并跳过重复的nums[i]
	//      III.当 s == 0; 记录 nums[k] + nums[i] + nums[j] ;i++, j-- 并跳过重复的i, j
	for k := 0; k < numLen-2; k++ {
		if nums[k] > 0 {
			break
		}
		if k-1 >= 0 && nums[k-1] == nums[k] {
			continue
		}
		i, j := k+1, numLen-1
		for i < j {
			s := nums[k] + nums[i] + nums[j]
			if s > 0 {
				j--
			} else {
				if s < 0 {
					i++
				} else {
					res = append(res, [3]int{nums[k], nums[i], nums[j]})
					for i < j && nums[i] == nums[i+1] {
						i++
					}
					for i < j && nums[j] == nums[j-1] {
						j--
					}
					i++
					j--
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}
