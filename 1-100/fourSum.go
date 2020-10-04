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

func fourSum2(nums []int, target int) [][]int {
	var ln = len(nums)
	if ln < 4 {
		return nil
	}
	// 先排序
	sort.Ints(nums)
	// 然后, emmm
	var i, j, k, l, r int
	var res [][]int
	r = target
	for i < ln-3 {
		r -= nums[i]
		for j = i + 1; j < ln-2; {
			r -= nums[j]
			k, l = j+1, ln-1
			for k < l {
				if v := nums[k] + nums[l]; v == r {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					for k = k + 1; k < l && nums[k] == nums[k-1]; k++ {
					}
					for l = l - 1; k < l && nums[l] == nums[l+1]; l-- {
					}
				} else if v > r {
					for l = l - 1; k < l && nums[l] == nums[l+1]; l-- {
					}
				} else {
					for k = k + 1; k < l && nums[k] == nums[k-1]; k++ {
					}
				}
			}
			r += nums[j]
			for j = j + 1; j < ln-2 && nums[j] == nums[j-1]; j++ {
			}
		}
		// 跳过相同的值
		for i = i + 1; i < ln-3 && nums[i] == nums[i-1]; i++ {
		}
		r = target
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

[1,-2,-5,-4,-3,3,3,5]
-11

*/

func main() {
	t := []int{1, -2, -5, -4, -3, 3, 3, 5}
	res := fourSum2(t, -11)
	fmt.Println(res)
}
