package main

import (
	"math"
	"sort"
)

func cross(p, q, r []int) int {
	return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
}

func getConvexHull(points [][]int) [][]int {
	n := len(points)
	if n < 4 {
		return points
	}

	// 按照 x 从小到大排序，如果 x 相同，则按照 y 从小到大排序
	sort.Slice(points, func(i, j int) bool {
		a, b := points[i], points[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] < b[1]
	})

	var hull [][]int
	// 求凸包的下半部分
	for _, p := range points {
		for len(hull) > 1 && cross(hull[len(hull)-2], hull[len(hull)-1], p) <= 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}
	// 求凸包的上半部分
	m := len(hull)
	for i := n - 1; i >= 0; i-- {
		for len(hull) > m && cross(hull[len(hull)-2], hull[len(hull)-1], points[i]) <= 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, points[i])
	}
	// hull[0] 同时参与凸包的上半部分检测，因此需去掉重复的 hull[0]
	return hull[:len(hull)-1]
}

func triangleArea(x1, y1, x2, y2, x3, y3 int) float64 {
	return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) / 2
}

func largestTriangleArea(points [][]int) (ans float64) {
	convexHull := getConvexHull(points)
	n := len(convexHull)
	for i, p := range convexHull {
		for j, k := i+1, i+2; j+1 < n; j++ {
			q := convexHull[j]
			for ; k+1 < n; k++ {
				curArea := triangleArea(p[0], p[1], q[0], q[1], convexHull[k][0], convexHull[k][1])
				nextArea := triangleArea(p[0], p[1], q[0], q[1], convexHull[k+1][0], convexHull[k+1][1])
				if curArea >= nextArea {
					break
				}
			}
			ans = math.Max(ans, triangleArea(p[0], p[1], q[0], q[1], convexHull[k][0], convexHull[k][1]))
		}
	}
	return
}
