package main

import (
	"fmt"
	"math"
)

func storeWater(bucket []int, vat []int) int {
	n := len(bucket)
	maxk := 0
	// 假定需要蓄水的此时
	for _, v := range vat {
		if v > maxk {
			maxk = v
		}
	}
	// 都为空时, 直接返回
	if maxk == 0 {
		return 0
	}

	res := math.MaxInt32
	for k := 1; k <= maxk && k < res; k++ {
		// k: 枚举蓄水的次数
		t := 0
		// t: 在使用这个次数的将vat[i]填满的情况下, 最少需要补几次水?
		for i := 0; i < n; i++ {
			// (vat[i] + k - 1) / k: 向上取整, 计算的就是k次蓄水至少需要 bucket[i] 中的水的数量
			// 再减去bucket[i]就是需要补的数量
			tt := (vat[i]+k-1)/k - bucket[i]
			t += max(0, tt)
			fmt.Println(k, vat[i], bucket[i], tt)
		}
		res = min(res, t+k)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
