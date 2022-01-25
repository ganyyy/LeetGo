package main

type Point map[int][]int

const (
	SHIFT = 10
	MASK  = (1 << SHIFT) - 1
)

func PackPos(x, y int) int {
	return x<<SHIFT | y
}

func UnpackPos(p int) (x, y int) {
	return p >> SHIFT, p & (MASK)
}

func (p Point) Add(k, v int) {
	p[k] = append(p[k], v)
}

type DetectSquares struct {
	count map[int]int
	xMap  Point
}

func Constructor() DetectSquares {
	return DetectSquares{
		count: make(map[int]int),
		xMap:  make(map[int][]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	var x, y = point[0], point[1]
	var p = PackPos(x, y)
	if _, ok := this.count[p]; !ok {
		this.xMap.Add(x, p)
	}
	this.count[p]++
}

func (this *DetectSquares) Count(point []int) int {
	var ret int

	var x, y = point[0], point[1]

	var xm = this.xMap[x]
	if len(xm) == 0 {
		return 0
	}

	for _, px := range xm {
		// xm存储的是x相同, y不同的点
		// 迭代所有x相同的点
		// 基于不同的y确定边长和另外的两个点
		var _, yy = UnpackPos(px)

		var cnt = this.count[px]

		var sub = y - yy
		if sub == 0 {
			continue
		}

		// 向左偏
		var p11, p12 = this.count[PackPos(x-sub, y)], this.count[PackPos(x-sub, yy)]
		if p11 != 0 && p12 != 0 {
			ret += p11 * p12 * cnt
		}

		// 向右偏
		p11, p12 = this.count[PackPos(x+sub, y)], this.count[PackPos(x+sub, yy)]
		if p11 != 0 && p12 != 0 {
			ret += p11 * p12 * cnt
		}

	}

	return ret
}
