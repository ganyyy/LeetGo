package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	// 先取第一个数字, 将其转化成三数求和的计算方式
	for p := 0; p < len(nums)-3; p++ {
		index := p
		value := nums[p]
		t := target - value
		// 如果和上一轮的值相同直接跳过
		if index > 0 && value == nums[index-1] {
			continue
		}
		for _i := index + 1; _i < len(nums)-2; _i++ {
			// 如果和上一轮的值相同直接跳过
			if _i > index+1 && nums[_i] == nums[_i-1] {
				continue
			}
			_j, _k := _i+1, len(nums)-1
			for _j < _k {
				sum := nums[_i] + nums[_j] + nums[_k]
				if sum == t {
					// 相等添加结果, 并且增加游标
					res = append(res, []int{value, nums[_i], nums[_j], nums[_k]})
					_j++
					for _j < _k && nums[_j] == nums[_j-1] {
						_j++
					}
					_k--
					for _j < _k && nums[_k] == nums[_k+1] {
						_k--
					}
				} else if sum < t {
					_j++
				} else {
					_k--
				}
			}
		}
	}
	return res
}

/**
[-3,-2,-1,0,0,1,2,3]
0

[0,1,5,0,1,5,5,-4]
11

[-1,0,1,2,-1,-4]
-1

[0,0,0,0]
0
*/

func main() {
	res := fourSum([]int{0, 0, 0, 0}, 0)
	fmt.Println(res)
}
