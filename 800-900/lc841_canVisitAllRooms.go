package main

func canVisitAllRooms(rooms [][]int) bool {
	// 标记是否去过某个房间
	m := make([]int, len(rooms))

	var helper func(i int)
	helper = func(i int) {
		if m[i] == 1 {
			return
		}
		m[i] = 1
		for _, v := range rooms[i] {
			helper(v)
		}
	}
	// 遍历一下位置, 如果存在为0的说明有房间走不进去
	helper(0)

	for _, v := range m {
		if v == 0 {
			return false
		}
	}
	return true
}
