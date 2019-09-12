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

func threeSumClosest(nums []int, target int) int {
	sLen := len(nums)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	getSum := func(k, i, j int) int {
		return nums[k] + nums[i] + nums[j]
	}
	sort.Ints(nums)
	min := 0
	minVal := 0
	init := true
	for k := 0; k < sLen-2; k++ {
		i, j := k+1, sLen-1
		for i < j {
			sum := getSum(k, i, j)
			dis := abs(sum - target)
			if dis == 0 {
				// 找到相同的就返回
				return sum
			}
			if init {
				minVal = sum
				min = dis
				init = false
			}
			if i+1 >= j {
				if dis < min {
					dis = min
					minVal = sum
				}
				break
			}
			vI, vJ := getSum(k, i+1, j), getSum(k, i, j-1)
			dI, dJ := abs(vI-target), abs(vJ-target)
			if dI == dJ {
				i++
				j--
			} else {
				if dI > dJ {
					if dJ < min {
						min = dJ
						minVal = vJ
					}
					j--
				} else {
					if dI < min {
						min = dI
						minVal = vI
					}
					i++
				}
			}
		}
	}
	return minVal
}

func main() {
	arrNums := [][]int{
		{1, 1, -1, -1, 3},
		{1, 2, 4, 8, 16, 32, 64, 128},
		{1, -3, 3, 5, 4, 1},
		{0, 1, 2},
		{0, 2, 1, -3},
		{1, 1, -1, -1, 3},
		{1, 1, 1, 0},
	}
	arrTarget := []int{
		-1,
		82,
		1,
		0,
		1,
		-1,
		100,
	}
	arrResult := []int{
		-1,
		82,
		1,
		3,
		0,
		-1,
		3,
	}

	for i := 0; i < len(arrResult); i++ {
		res := threeSumClosest(arrNums[i], arrTarget[i])
		if res != arrResult[i] {
			fmt.Printf("i:%d\nnums:%v, \n target:%d, expect:%d, result:%d\n", i, arrNums[i], arrTarget[i], arrResult[i], res)
			fmt.Printf("-------------------\n")
		}
	}
	//res := threeSumClosest(arrNums[4], arrTarget[4])
	//fmt.Printf("%d\n",res)
	//if res == arrResult[4] {
	//}
}
