package main

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	count := make(map[int]int)
	for i := 0; i < len(deck); i++ {
		count[deck[i]]++
	}
	var x int
	// 保证所有牌的数量间的最大公约数不为1
	for _, v := range count {
		if v == 0 {
			continue
		}
		if x == 0 {
			x = v
		} else {
			x = gcd(x, v)
			if x == 1 {
				return false
			}
		}
	}
	return true
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {

}
