package main

import "fmt"

func subarraysWithKDistinct(A []int, K int) int {
	return subarraysWithMostK(A, K) - subarraysWithMostK(A, K-1)
}

// 这个函数是求子数组中不同的数字最多有K个 的 子数组的个数
func subarraysWithMostK(A []int, K int) int {
	var set = make(map[int]int, len(A))
	var l, r, res int
	for r < len(A) {
		set[A[r]]++
		r++
		for len(set) > K {
			set[A[l]]--
			if set[A[l]] == 0 {
				delete(set, A[l])
			}
			l++
		}
		res += r - l + 1
	}
	return res
}

func subarraysWithKDistinct2(A []int, K int) int {
	if A == nil || len(A) < K {
		return 0
	}

	hash := make([]int, len(A)+1)
	l, count, result, results := 0, 0, 1, 0
	for r := 0; r < len(A); r++ {
		hash[A[r]]++

		if hash[A[r]] == 1 {
			count++
		}

		// 这一步很巧妙的地方就是
		// 相同的数字提前计算了 结果的个数
		for hash[A[l]] > 1 || count > K {
			if count > K {
				result = 1
				count--
			} else {
				result++
			}
			hash[A[l]]--
			l++
		}

		if count == K {
			results += result
		}
	}

	return results
}

func main() {
	/*
		[2,1,1,1,2]
		1
	*/
	fmt.Println(subarraysWithKDistinct([]int{1, 1, 1}, 1))
}
