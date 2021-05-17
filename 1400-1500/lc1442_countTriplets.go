package main

func countTriplets(arr []int) int {
	// 暴力枚举
	// 前缀和
	var sum = make([]int, len(arr)+1)
	copy(sum[1:], arr)

	for i := 1; i < len(sum); i++ {
		sum[i] ^= sum[i-1]
	}

	var count int

	for i := 1; i < len(sum)-1; i++ {
		for j := i + 1; j < len(sum); j++ {
			for k := j; k < len(sum); k++ {
				if sum[i-1]^sum[j-1] == sum[j-1]^sum[k] {
					count++
				}
			}
		}
	}
	return count

}

func countTripletsBig(arr []int) int {
	var count int
	var sum int
	for i := 0; i < len(arr)-1; i++ {
		sum = 0
		for k := i; k < len(arr); k++ {
			sum ^= arr[k]
			// 把j想象成(i,k)中的某个值
			// 因为 [i,k]的异或和为0, 说明在[i,k]中间必有一个j使得[i,j-1] == [j,k]
			if sum == 0 && k > i {
				count += k - i
			}
		}
	}

	return count
}
