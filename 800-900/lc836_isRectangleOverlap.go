package main

import "fmt"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 判断一下, x远的是rec2，近的是rec1
	//if rec1[2] > rec2[2] {
	//	rec1, rec2 = rec2, rec1
	//}
	// 首先必须满足 x相交, 其次满足y相交
	//return rec1[2] > rec2[0] && ((rec2[3] >= rec1[3] && rec1[3] > rec2[1]) || (rec1[3] >= rec2[3] && rec2[3] > rec1[1]))

	// 返回不重叠
	return !(rec1[2] <= rec2[0] || // left
		rec1[3] <= rec2[1] || // bottom
		rec1[0] >= rec2[2] || // right
		rec1[1] >= rec2[3]) // top

}

func main() {
	/**
	[2,17,6,20]
	[3,8,6,20]
	*/
	fmt.Println(isRectangleOverlap([]int{2, 17, 6, 20}, []int{3, 8, 6, 20}))
}
