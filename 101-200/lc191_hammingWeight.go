package main

func hammingWeight(num uint32) int {
	var count int
	// 核心想法是通过-1将低位的1变成0
	// 如 0100 0011
	for num > 0 {
		num &= num - 1
		count++
	}
	return count
}

func main() {

}
