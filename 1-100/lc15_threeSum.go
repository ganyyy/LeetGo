package main

import (
	"fmt"
	"math"
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
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("error test:%v, %d\n", nums, target)
		}
	}()
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
	minVal := nums[0] + nums[1] + nums[2]
	minDis := abs(minVal - target)
	for k := 0; k < sLen-2; k++ {
		i, j := k+1, sLen-1
		for i < j {
			sum := getSum(k, i, j)
			dis := abs(sum - target)
			if dis == 0 {
				// 找到相同的就返回
				return sum
			}
			if minDis > dis {
				minDis = dis
				minVal = sum
			}
			if sum > target {
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				j--
			} else {
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				i++
			}
		}
	}
	return minVal
}

func threeSumClosest2(nums []int, target int) int {
	// 先进行排序
	sort.Ints(nums)
	res := math.MaxInt32
	cur := abs(res, target)
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}

			if t := abs(target, sum); cur > t {
				cur = t
				res = sum
			}

			if sum > target {
				for ; k > j && nums[k] == nums[k-1]; k-- {
				}
				k--
			} else {
				for ; j < k && nums[j] == nums[j+1]; j++ {
				}
				j++
			}
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	// arrNums := [][]int{
	//	{1, 1, -1, -1, 3},
	//	{1, 2, 4, 8, 16, 32, 64, 128},
	//	{1, -3, 3, 5, 4, 1},
	//	{0, 1, 2},
	//	{0, 2, 1, -3},
	//	{1, 1, -1, -1, 3},
	//	{1, 1, 1, 0},
	//	{13, 2, 0, -14, -20, 19, 8, -5, -13, -3, 20, 15, 20, 5, 13, 14, -17, -7, 12, -6, 0, 20, -19, -1, -15, -2, 8, -2, -9, 13, 0, -3, -18, -9, -9, -19, 17, -14, -19, -4, -16, 2, 0, 9, 5, -7, -4, 20, 18, 9, 0, 12, -1, 10, -17, -11, 16, -13, -14, -3, 0, 2, -18, 2, 8, 20, -15, 3, -13, -12, -2, -19, 11, 11, -10, 1, 1, -10, -2, 12, 0, 17, -19, -7, 8, -19, -17, 5, -5, -10, 8, 0, -12, 4, 19, 2, 0, 12, 14, -9, 15, 7, 0, -16, -5, 16, -12, 0, 2, -16, 14, 18, 12, 13, 5, 0, 5, 6},
	// }
	// arrTarget := []int{
	//	-1,
	//	82,
	//	1,
	//	0,
	//	1,
	//	-1,
	//	100,
	//	-59,
	// }
	// arrResult := []int{
	//	-1,
	//	82,
	//	1,
	//	3,
	//	0,
	//	-1,
	//	3,
	//	-58,
	// }
	//
	// for i := 0; i < len(arrResult); i++ {
	//	res := threeSumClosest(arrNums[i], arrTarget[i])
	//	if res != arrResult[i] {
	//		fmt.Printf("i:%d\nnums:%v, \n target:%d, expect:%d, result:%d\n", i, arrNums[i], arrTarget[i], arrResult[i], res)
	//		fmt.Printf("-------------------\n")
	//	}
	// }

	fmt.Println(threeSumClosest2([]int{-1, 2, 1, -4}, 1))

	/*
		[]
		-59
	*/
	// res := threeSumClosest(arrNums[6], arrTarget[6])
	// fmt.Printf("%d\n",res)
	// if res == arrResult[4] {
	// }
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)

	// [i:n-2]
	// [i+1:n-1]
	var ret [][]int
	numsLength := len(nums)
	for i := 0; i < numsLength-2; i++ {
		numI := nums[i]
		if numI > 0 {
			break
		}
		if i > 0 && nums[i-1] == numI {
			continue
		}
		for j, k := i+1, numsLength-1; j < k; {
			numJ, numK := nums[j], nums[k]
			sum := numI + numJ + numK
			var subK, addJ bool
			if sum > 0 {
				subK = true
			} else if sum < 0 {
				addJ = true
			} else {
				ret = append(ret, []int{numI, numJ, numK})
				subK = true
				addJ = true
			}
			if subK {
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				}
			}
			if addJ {
				for j++; k > j && nums[j] == nums[j-1]; j++ {
				}
			}
		}
	}
	return ret
}
