package main

const MAX = 1e9

var Map map[[10]int]bool

func init() {
	Map = make(map[[10]int]bool)
	for i := 1; i < MAX; i <<= 1 {
		Map[ParseNum(i)] = true
	}
}

func ParseNum(n int) (val [10]int) {
	for j := n; j != 0; j /= 10 {
		val[j%10]++
	}
	return
}

func reorderedPowerOf2(n int) bool {
	return Map[ParseNum(n)]
}
