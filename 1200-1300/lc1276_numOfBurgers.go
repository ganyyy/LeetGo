package main

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	// 鸡兔同笼是吧?
	/*
	   4x+2y = tomatoSlices ①
	   x+y = cheeseSlices   ②

	   ② * 2 => 2x+2y = 2*cheeseSlices ③
	   ① - ③ => 2x = tomatoSlices - 2*cheeseSlices ④

	   x = ④ / 2
	   y = ② - x
	*/
	doubleX := tomatoSlices - 2*cheeseSlices
	if doubleX < 0 || doubleX%2 != 0 {
		return nil
	}
	x := doubleX / 2
	y := cheeseSlices - x
	if y < 0 {
		return nil
	}
	return []int{x, y}
}
