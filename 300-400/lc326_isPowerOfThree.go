package main

func isPowerOfThree(n int) bool {
	// 核心: 3在int范围内的最大次幂是19
	// 质数的整数次幂只有除其整数次幂结果才会是一个整数
	// 如果取余结果为0, 说明n是整数次幂
	return n > 0 && 1162261467%n == 0
}

func main() {

}
