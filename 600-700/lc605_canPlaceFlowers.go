package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if len(flowerbed) == 1 {
		return flowerbed[0] == 0
	}
	// 必须一半以上才能存的下去
	if (len(flowerbed)+1)>>1 < n {
		return false
	}
	var m = len(flowerbed) - 1
	for i := 0; i <= m; i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if i > 0 && flowerbed[i-1] != 0 {
			continue
		}
		if i < m && flowerbed[i+1] != 0 {
			continue
		}
		// 这里能插入一朵花
		flowerbed[i] = 1
		n--
		if n == 0 {
			return true
		}
	}

	return n == 0
}

func canPlaceFlowers2(flowerbed []int, n int) bool {
	lenth := len(flowerbed)
	for index := 0; index < lenth && n > 0; index += 2 {
		if flowerbed[index] == 0 {
			if index+1 == lenth || flowerbed[index+1] == 0 {
				n--
			} else {
				index++
			}
		}
	}
	return n == 0
}

func main() {
	println(canPlaceFlowers([]int{0, 0, 0}, 2))
}
