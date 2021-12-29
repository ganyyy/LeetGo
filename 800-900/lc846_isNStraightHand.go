package main

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	if groupSize == 1 {
		return true
	}

	sort.Ints(hand)
	for i := 0; i < len(hand); i++ {
		// 如果已经被选取了, 直接跳过
		if hand[i] == -1 {
			continue
		}

		// 以当前数字为基准, 选取连续的数组. 使其个数等同于groupSize
		value, cnt, isFull := hand[i], 0, false
		for j := i; j < len(hand); j++ {
			// 被选取的直接跳过
			if hand[j] == -1 {
				continue
			}

			// 合适的就选一下
			if hand[j] == value+cnt {
				hand[j] = -1
				cnt++
				if cnt == groupSize {
					isFull = true
					break
				}
			}

			// 如果满足分组需求的话, 必定会保证 value+groupSize在次之前出现并跳出循环
			if hand[j] >= value+groupSize {
				return false
			}
		}

		// 避免没有完成迭代的情况
		if !isFull {
			return false
		}
	}
	return true
}

func isNStraightHand2(hand []int, groupSize int) bool {
	if len(hand)%groupSize > 0 {
		return false
	}
	sort.Ints(hand)
	cnt := map[int]int{}
	for _, num := range hand {
		cnt[num]++
	}
	for _, x := range hand {
		if cnt[x] == 0 {
			continue
		}
		for num := x; num < x+groupSize; num++ {
			if cnt[num] == 0 {
				return false
			}
			cnt[num]--
		}
	}
	return true
}

func isNStraightHand3(hand []int, groupSize int) bool {

	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)

	var min, max = hand[0], hand[len(hand)-1]

	var cnt = make(map[int]int)
	for _, v := range hand {
		cnt[v]++
	}

	for i := min; i <= max-groupSize+1; i++ {
		var old, ok = cnt[i]
		if !ok {
			continue
		}
		// 全部消耗
		delete(cnt, i)
		for j := i + 1; j < i+groupSize; j++ {
			v, ok := cnt[j]
			if !ok || v < old {
				return false
			} else if v-old == 0 {
				delete(cnt, j)
			} else {
				cnt[j] = v - old
			}
		}
	}

	return len(cnt) == 0
}
