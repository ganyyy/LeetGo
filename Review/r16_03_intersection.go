package main

// 看的头疼, 不做了..
func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	x1, y1, x2, y2, x3, y3, x4, y4 := start1[0], start1[1], end1[0], end1[1], start2[0], start2[1], end2[0], end2[1]
	var det = func(a, b, c, d int) int {
		return a*d - b*c
	}

	d := det(x1-x2, x4-x3, y1-y2, y4-y3)
	p := det(x4-x2, x4-x3, y4-y2, y4-y3)
	q := det(x1-x2, x4-x2, y1-y2, y4-y2)

	if d != 0 {
		lam, eta := p/d, q/d
		if lam < 0 || lam > 1 || eta < 0 || eta > 1 {
			return []float64{}
		}
		return []float64{float64(lam*x1 + (1-lam)*x2), float64(lam*y1 + (1-lam)*y2)}
	}
	if p != 0 || q != 0 {
		return []float64{}
	}

	return []float64{}
}

func main() {

}
