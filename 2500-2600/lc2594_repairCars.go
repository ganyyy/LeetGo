package main

import (
	"math"
	"sort"
)

func repairCars(ranks []int, cars int) int64 {
	// 计数
	var allCnt [101]int
	// rank最小值
	var minR = math.MaxInt32
	for _, rank := range ranks {
		if rank < minR {
			minR = rank
		}
		allCnt[rank]++
	}
	const (
		Shift = 32
		Mask  = (1 << Shift) - 1
	)
	// 压缩
	var rankCntBuf = allCnt[:0]
	for rank, cnt := range allCnt {
		if cnt == 0 {
			continue
		}
		rankCntBuf = append(rankCntBuf, (rank<<Shift)|cnt)
	}
	// 上界就是能力最强的人(rank最低)修所有的车, 因为只要有其他人加入, 那么时间就会减少
	return int64(sort.Search(minR*cars*cars, func(curTime int) bool {
		s := 0
		for _, rankCnt := range rankCntBuf {
			rank := rankCnt >> Shift
			cnt := rankCnt & Mask
			s += int(math.Sqrt(float64(curTime/rank))) * cnt
		}
		// 累计每个工人的修车数, 二分查找答案.
		return s >= cars
	}))
}
