package main

import "sort"

func fullBloomFlowers(flowers [][]int, people []int) []int {
	n := len(flowers)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, f := range flowers {
		starts[i] = f[0]
		ends[i] = f[1]
	}
	// 开花和凋落的时间分别排序
	sort.Ints(starts)
	sort.Ints(ends)

	for i, p := range people {
		// 注意:
		// 在start中查找p+1是查找的 <= p位置的数所处的位置
		// 在 ends中查找 p 是查找的  < p位置的数所处的位置
		people[i] = sort.SearchInts(starts, p+1) - sort.SearchInts(ends, p)
	}
	return people
}

func fullBloomFlowers2(flowers [][]int, people []int) []int {
	// 构建离散化差分数组
	var flowerDiff = make(map[int]int, len(flowers))
	for _, flower := range flowers {
		flowerDiff[flower[0]]++
		flowerDiff[flower[1]+1]--
	}
	length := len(flowerDiff)
	// 按照增删的关键时间点排序
	var flowerTimes = make([]int, 0, length)
	for t := range flowerDiff {
		flowerTimes = append(flowerTimes, t)
	}
	sort.Ints(flowerTimes)

	// 先构建用户id映射, 再将用户按照观赏时间排序
	var peopleId = make([]int, len(people))
	for id := range people {
		peopleId[id] = id
	}
	sort.Slice(peopleId, func(i, j int) bool { return people[peopleId[i]] < people[peopleId[j]] })

	// 此时的玩家参观的顺序, 以及花开花落的关键时间点都是时间正序
	// 可以使用双指针进行迭代
	// curFlower: 到某个用户观赏时间点时的累加和
	// flowerIdx: 迭代flowerTimes的位置
	var flowerIdx, curFlower int
	for _, id := range peopleId {
		peopleEnd := people[id]
		for ; flowerIdx < length && flowerTimes[flowerIdx] <= peopleEnd; flowerIdx++ {
			curFlower += flowerDiff[flowerTimes[flowerIdx]]
		}
		people[id] = curFlower
	}
	return people
}
