package main

import "math"

func isRectangleCover(rectangles [][]int) bool {

	// 最左边(minx), 最右边(maxx), 最上边(miny), 最下边(maxy)
	// 计算面积, 对比每个矩形的面积
	var totalArea int

	var set = make(map[int]struct{}, 16)

	var addToSet = func(v [2]int) {
		var val = v[0]*1e8 + v[1]
		if _, ok := set[val]; ok {
			delete(set, val)
		} else {
			set[val] = struct{}{}
		}
	}

	var checkInSet = func(v [2]int) bool {
		var val = v[0]*1e8 + v[1]
		_, ok := set[val]
		return ok
	}

	var Min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var Max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var mix, max, miy, may = math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32

	for _, r := range rectangles {
		totalArea += GetArea(r)

		// 完美矩形的前提是, 4个点仅出现一次, 其余的点成对出现!

		var a, b, c, d = [2]int{r[0], r[1]}, [2]int{r[2], r[1]}, [2]int{r[2], r[3]}, [2]int{r[0], r[3]}

		// 获取四个顶点, 用于计算总面积
		mix = Min(mix, r[0])
		miy = Min(miy, r[1])
		max = Max(max, r[2])
		may = Max(may, r[3])

		addToSet(a)
		addToSet(b)
		addToSet(c)
		addToSet(d)
	}

	if len(set) != 4 {
		return false
	}

	if !checkInSet([2]int{mix, miy}) ||
		!checkInSet([2]int{max, may}) ||
		!checkInSet([2]int{max, miy}) ||
		!checkInSet([2]int{mix, may}) {
		return false
	}

	return (may-miy)*(max-mix) == totalArea

}

func GetArea(rectangle []int) int {
	return (rectangle[2] - rectangle[0]) * (rectangle[3] - rectangle[1])
}
