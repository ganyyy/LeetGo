package main

func add(a int, b int) int {
	for (a & b) != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a ^ b
}

func main() {

}
