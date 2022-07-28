package main

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {

	var distance = func(p1, p2 []int) int {
		var a, b = p1[0] - p2[0], p1[1] - p2[1]
		return a*a + b*b
	}

	var set = make(map[int]struct{})
	var empty struct{}

	// 正负两个值, 且不能为0
	set[distance(p1, p2)] = empty
	set[distance(p1, p3)] = empty
	set[distance(p1, p4)] = empty
	set[distance(p2, p3)] = empty
	set[distance(p2, p4)] = empty
	set[distance(p3, p4)] = empty

	// 如果出现了 [0,0],[1,1],[0,0],[1,1]这种情况, 也会爆炸, 因为重合点
	// 所以需要判定不存在重合点
	_, ok := set[0]
	return len(set) == 2 && !ok
}
